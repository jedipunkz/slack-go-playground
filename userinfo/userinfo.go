package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/mitchellh/go-homedir"
	"github.com/slack-go/slack"
	"github.com/spf13/viper"
)

const (
	hookString = "userslimit"
)

var (
	botID string
)

// Bot is struct for bot
type Bot struct {
	api *slack.Client
	rtm *slack.RTM
}

// NewBot is function for bot
func NewBot(token string) *Bot {
	bot := new(Bot)
	bot.api = slack.New(token)
	bot.rtm = bot.api.NewRTM()
	return bot
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
	token := viper.GetString("token")

	api := slack.New(token)

	// some, _ := api.GetTeamInfo()
	some, _ := api.GetUserInfo("US1537YGM")

	//
	// groups, err := api.GetGroups(false)
	// if err != nil {
	// 	fmt.Printf("%s\n", err)
	// }

	// var (
	// 	postAsUserName string
	// 	postAsUserID   string
	// 	postToUserName string
	// 	postToUserID   string
	// )

	// authTest, err := api.AuthTest()
	// if err != nil {
	// 	fmt.Printf("Error getting channels: %s\n", err)
	// 	return
	// }

	// // Post as the authenticated user.
	// postAsUserName = authTest.User
	// postAsUserID = authTest.UserID
	// // // Posting to DM with self causes a conversation with slackbot.
	// postToUserName = authTest.User
	// postToUserID = authTest.UserID

	bot := NewBot(token)
	go bot.rtm.ManageConnection()

	for {
		select {
		case msg := <-bot.rtm.IncomingEvents:

			switch ev := msg.Data.(type) {

			case *slack.ConnectedEvent:
				botID = ev.Info.User.ID

			case *slack.MessageEvent:
				// user := ev.User
				// userInfo, _ := bot.api.GetUserInfo(ev.User)
				// user := userInfo.Profile.RealName
				text := ev.Text
				shellSlack := strings.Replace(text, "<@"+botID+">", "", 1)
				shellSlack = strings.TrimSpace(shellSlack)

				if ev.Type == "message" && strings.HasPrefix(shellSlack, hookString) && strings.HasPrefix(text, "<@"+botID+">") {
					bot.rtm.SendMessage(bot.rtm.NewOutgoingMessage("some: "+some.Name+" "+some.TeamID, ev.Channel))
					// bot.rtm.SendMessage(bot.rtm.NewOutgoingMessage("postAsUserName: "+postAsUserName+"です！！！１", ev.Channel))
					// bot.rtm.SendMessage(bot.rtm.NewOutgoingMessage("postAsUserID: "+postAsUserID+"です！！！１", ev.Channel))
					// bot.rtm.SendMessage(bot.rtm.NewOutgoingMessage("postToUserName: "+postToUserName+"です！！！１", ev.Channel))
					// bot.rtm.SendMessage(bot.rtm.NewOutgoingMessage("postToUserID: "+postToUserID+"です！！！１", ev.Channel))
				}
			}
		}
	}
}
