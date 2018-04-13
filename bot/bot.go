package bot

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-redis/redis"
	"github.com/markwallsgrove/hercules/workers"
	"github.com/nlopes/slack"
)

// Bot slack bot
type Bot struct {
	listeners      []workers.Registration
	workers        []workers.Worker
	incomingEvents *chan slack.RTMEvent
	quitSignal     *chan os.Signal
	rtm            *slack.RTM
	logger         *log.Logger
	memory         *redis.Client
}

// Register a worker to recieve events
func (b *Bot) Register(worker workers.Worker) {
	b.workers = append(b.workers, worker)
	listeners := worker.Init(b.rtm, b.memory)
	b.listeners = append(b.listeners, listeners...)
}

// Listen to events from slack. This call blocks.
func (b *Bot) Listen() {
	for {
		select {
		case event := <-*b.incomingEvents:
			b.processEvent(event)
			break
		case signal := <-*b.quitSignal:
			b.logger.Printf("Recieved signal: %v", signal)
			break
		}
	}

	b.logger.Println("Shutting down workers")
	for _, worker := range b.workers {
		worker.Quit()
	}

	b.logger.Println("Closing down connection to slack")
	if err := b.rtm.Disconnect(); err != nil {
		b.logger.Fatal("Error disconnecting from slack", err)
	}
}

func (b *Bot) processEvent(event slack.RTMEvent) {
	b.logger.Println("processing message ", event.Data)

	switch event.Data.(type) {
	case slack.MessageEvent:
		b.logger.Println("checking string '", event.Data.(slack.MessageEvent).Text)
		b.processMessage(event.Data.(slack.MessageEvent))
		break
	}
}

func (b *Bot) processMessage(event slack.MessageEvent) {
	for _, listener := range b.listeners {
		// TODO: mddleware
		listener.Apply(event)
	}
}

// Quit listening to slack. This will cause the listen method to quit.
func (b *Bot) Quit() {
	b.logger.Println("Requesting shutdown")
	*b.quitSignal <- syscall.SIGTERM
}

// MakeBot create a new bot
func MakeBot(logger *log.Logger, apiSecret string) (*Bot, error) {
	api := slack.New(apiSecret)
	api.SetDebug(true)

	slack.SetLogger(logger)

	rtm := api.NewRTM()
	go rtm.ManageConnection()

	// TODO: move address & db to argument
	memory := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	if _, err := memory.Ping().Result(); err != nil {
		return nil, err
	}

	quitChannel := make(chan os.Signal)
	signal.Notify(quitChannel, syscall.SIGTERM)
	signal.Notify(quitChannel, syscall.SIGINT)

	return &Bot{
		incomingEvents: &rtm.IncomingEvents,
		quitSignal:     &quitChannel,
		logger:         logger,
		rtm:            rtm,
		memory:         memory,
	}, nil
}
