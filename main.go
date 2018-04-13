package main

import (
	"log"
	"os"

	"github.com/markwallsgrove/hercules/bot"
	"github.com/markwallsgrove/hercules/workers"
)

func main() {
	slackSecret := os.Getenv("SLACK_API_TOKEN")
	logger := log.New(os.Stdout, "bot: ", log.Lshortfile|log.LstdFlags)

	if bot, err := bot.MakeBot(logger, slackSecret); err != nil {
		panic(err)
	} else {
		bot.Register(workers.MakeTestWorker(logger))
		bot.Listen()
		bot.Quit()
	}
}
