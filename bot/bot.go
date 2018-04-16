package bot

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/markwallsgrove/hercules/types"
	"github.com/markwallsgrove/hercules/workers"
	"github.com/nlopes/slack"
)

// Bot slack bot
type Bot struct {
	listeners      []workers.Registration
	workers        []workers.Worker
	incomingEvents *chan slack.RTMEvent
	quitSignal     *chan os.Signal
	rtm            types.RTM
	logger         *log.Logger
}

// Register a worker to recieve events
func (b *Bot) Register(worker workers.Worker) {
	listeners := worker.Init(b.rtm)
	b.listeners = append(b.listeners, listeners...)
	b.workers = append(b.workers, worker)
}

// Listen to events from slack. This call blocks.
func (b *Bot) Listen() {
	defer b.close()

	for {
		select {
		case event := <-*b.incomingEvents:
			b.processEvent(event)
		case signal := <-*b.quitSignal:
			b.logger.Printf("recieved signal: %v", signal)
			return
		}
	}
}

// Process incoming slack event
func (b *Bot) processEvent(event slack.RTMEvent) {
	b.logger.Println("processing message ", event.Data)

	switch event.Data.(type) {
	case *slack.MessageEvent:
		messageEvent := event.Data.(*slack.MessageEvent)
		b.processMessage(messageEvent)
		break
	}
}

// Process incoming message event
func (b *Bot) processMessage(event *slack.MessageEvent) {
	for _, listener := range b.listeners {
		listener.Apply(event)
	}
}

// close down all resources
func (b *Bot) close() {
	b.logger.Println("shutting down workers")
	for _, worker := range b.workers {
		worker.Quit()
	}

	b.logger.Println("closing down connection to slack")
	if err := b.rtm.Disconnect(); err != nil {
		b.logger.Fatal("disconnecting from slack", err)
	}

	b.logger.Println("closing down queues")
	close(*b.quitSignal)
}

// Quit listening to slack. This will cause the listen method to quit.
func (b *Bot) Quit() {
	b.logger.Println("Requesting shutdown")
	*b.quitSignal <- syscall.SIGTERM
}

func StartBot(secret string, url string) *Bot {
	logger := log.New(os.Stdout, "bot: ", log.Lshortfile|log.LstdFlags)

	slack.SetLogger(logger)
	slack.SLACK_API = url

	api := slack.New(secret)
	api.SetDebug(true)

	rtm := api.NewRTM(slack.RTMOptionUseStart(true))
	go rtm.ManageConnection()

	quitChannel := make(chan os.Signal)
	signal.Notify(quitChannel, syscall.SIGTERM)
	signal.Notify(quitChannel, syscall.SIGINT)

	bot := &Bot{
		incomingEvents: &rtm.IncomingEvents,
		quitSignal:     &quitChannel,
		logger:         logger,
		rtm:            rtm,
	}

	bot.Register(workers.MakeTestWorker(logger))
	return bot
}
