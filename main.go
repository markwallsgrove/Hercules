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
	bot := bot.MakeBot(logger, slackSecret)
	bot.Register(workers.MakeTestWorker(logger))
	bot.Register(workers.MakeOutputWorker(logger))
	bot.Listen()
	bot.Quit()
}
