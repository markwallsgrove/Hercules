package main

import (
	"log"
	"os"

	"github.com/markwallsgrove/hercules/bot"
	"github.com/markwallsgrove/hercules/workers"
)

func main() {
	logger := log.New(os.Stdout, "bot: ", log.Lshortfile|log.LstdFlags)
	bot := bot.MakeBot(logger, "")
	bot.Register(workers.MakeTestWorker(logger))
	bot.Register(workers.MakeOutputWorker(logger))
	bot.Listen()
	bot.Quit()
}
