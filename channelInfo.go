package main

import (
	"strings"

	"github.com/nlopes/slack"
)

type chanInfo struct {
	private bool
}

type channelInfo struct {
	info map[string]chanInfo
}

var (
	info = make(map[string]chanInfo)
)

func CreateChannelInfo(channel *slack.Channel) chanInfo {
	return chanInfo{
		private: channel.IsIM || channel.IsGroup,
	}
}

func CreateDefaultChannelInfo(channelID string) chanInfo {
	isPrivate := strings.HasPrefix(channelID, "D")

	return chanInfo{
		private: isPrivate,
	}
}

func ChannelIsPrivate(channelID string) bool {
	chanInfo, ok := info[channelID]
	if !ok {
		channel, err := Config().SlackAPI().GetChannelInfo(channelID)
		if err != nil {
			return true
		}

		chanInfo = CreateChannelInfo(channel)

		info[channelID] = chanInfo
	}

	return chanInfo.private
}
