package main

import (
	"fmt"
	"runtime"

	"github.com/nlopes/slack"
)

var queue string

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

		case *slack.ConnectedEvent:
			handleConnection(msg.Data.(*slack.ConnectedEvent))

		case *slack.InvalidAuthEvent:
			fmt.Printf("Invalid credentials\n")
			return

		default:
		}
	}

}
