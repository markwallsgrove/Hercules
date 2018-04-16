package main

import (
	"log"
	"os"

	"github.com/markwallsgrove/hercules/bot"
	"github.com/markwallsgrove/hercules/workers"
	"github.com/nlopes/slack"
)

func main() {
	logger := log.New(os.Stdout, "bot: ", log.Lshortfile|log.LstdFlags)
	secret := os.Getenv("SLACK_API_TOKEN")
	bot := bot.StartBot(secret, slack.SLACK_API)
	bot.Register(workers.MakeTestWorker(logger))
	bot.Listen()
}
