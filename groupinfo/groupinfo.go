package main

import (
	"fmt"
	"os"

	"github.com/mitchellh/go-homedir"
	"github.com/slack-go/slack"
	"github.com/spf13/viper"
)

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
	token := viper.GetString("token")

	api := slack.New(token)

	// 権限が無いのか、空のスライスが帰ってくる
	groups, _ := api.GetUserGroups()
	fmt.Println(groups)
}
