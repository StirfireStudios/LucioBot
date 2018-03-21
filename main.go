package main

import (
	"fmt"
	"runtime"

	"github.com/nlopes/slack"
)

var queue string

func handleMessage(msg *slack.MessageEvent) {
	if msg.Type != "message" {
		return
	}

	if msg.Hidden {
		return
	}

	if len(msg.Text) == 0 {
		return
	}

	private := ChannelIsPrivate(msg.Channel)
	_ = private

	fmt.Printf("Message: %s\n", msg.Text)
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	if !Config().Load("config.json") {
		fmt.Println("Unable to parse/load config!")
		return
	}

	rtm := Config().SlackAPI().NewRTM()
	go rtm.ManageConnection()
	for msg := range rtm.IncomingEvents {
		switch msg.Data.(type) {
		case *slack.MessageEvent:
			handleMessage(msg.Data.(*slack.MessageEvent))

		case *slack.InvalidAuthEvent:
			fmt.Printf("Invalid credentials\n")
			return

		default:
		}
	}

}
