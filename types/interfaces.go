package types

import "github.com/nlopes/slack"

type RTM interface {
	SendMessage(msg *slack.OutgoingMessage)
	NewOutgoingMessage(text string, channelID string) *slack.OutgoingMessage
	Disconnect() error
}
