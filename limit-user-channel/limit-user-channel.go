package main

import (
	"fmt"
	"os"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

// Users is struct for users
type Users []struct {
	Name string
	ID   string
}

func init() {
	home, err := homedir.Dir()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	viper.AddConfigPath(home)
	viper.SetConfigName(".infra-bot")

	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file: ", viper.ConfigFileUsed())
	} else {
		fmt.Println("Could not read config file: ", err)
		os.Exit(1)
	}
}

func main() {
	if getExecLimitUser("US1537YGM") {
		fmt.Println("ok")
	} else {
		fmt.Println("not ok")
	}

	if getExecLimitChannel("xixixi") {
		fmt.Println("ok")
	} else {
		fmt.Println("not ok")
	}
}

func getExecLimitUser(slackUserID string) bool {
	var users Users

	viper.UnmarshalKey("samplebot.limits.users", &users)

	for _, user := range users {
		if slackUserID == user.ID {
			return true
		}
	}

	return false
}

func getExecLimitChannel(channelID string) bool {
	allowedChannels := viper.GetStringSlice("samplebot.limits.channel")

	for _, allowedChannel := range allowedChannels {
		if channelID == allowedChannel {
			return true
		}
	}

	return false
}
