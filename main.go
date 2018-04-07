package main

import (
	"log"
	"os"

	"github.com/markwallsgrove/hercules/bot"
	"github.com/markwallsgrove/hercules/workers"
)

func main() {
	logger := log.New(os.Stdout, "bot: ", log.Lshortfile|log.LstdFlags)
	bot := bot.MakeBot(logger, "xoxb-329560334997-dhVE3F5IMMFP2u8Kz9sFOubu")
	bot.Register(workers.MakeTestWorker(logger))
	bot.Register(workers.MakeOutputWorker(logger))
	bot.Listen()
	bot.Quit()
}
