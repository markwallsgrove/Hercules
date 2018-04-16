package main

import (
	"os"

	"github.com/markwallsgrove/hercules/bot"
	"github.com/nlopes/slack"
)

func main() {
	secret := os.Getenv("SLACK_API_TOKEN")
	bot := bot.StartBot(secret, slack.SLACK_API)
	bot.Listen()
}
