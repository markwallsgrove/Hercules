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
	workers    []workers.Worker
	quitSignal chan os.Signal
	rtm        *slack.RTM
	logger     *log.Logger
	memory     *redis.Client
}

// Register a worker to recieve events
func (b *Bot) Register(worker workers.Worker) {
	b.workers = append(b.workers, worker)
	worker.Init(b.rtm)
}

// Listen to events from slack. This call blocks.
func (b *Bot) Listen() {
	b.quitSignal = make(chan os.Signal)
	signal.Notify(b.quitSignal, syscall.SIGTERM)
	signal.Notify(b.quitSignal, syscall.SIGINT)

	for {
		select {
		case event := <-b.rtm.IncomingEvents:
			for _, worker := range b.workers {
				if worker.Process(event) {
					break
				}
			}
		case signal := <-b.quitSignal:
			b.logger.Printf("Recieved signal: %v", signal)
			return
		}
	}

}

// Quit listening to slack. This will cause the listen method to quit.
func (b *Bot) Quit() {
	b.logger.Println("Requesting shutdown")
	b.quitSignal <- syscall.SIGTERM

	b.logger.Println("Shutting down workers")
	for _, worker := range b.workers {
		worker.Quit()
	}

	b.logger.Println("Closing down connection to slack")
	if err := b.rtm.Disconnect(); err != nil {
		b.logger.Fatal("Error disconnecting from slack", err)
	}
}

// MakeBot create a new bot
func MakeBot(logger *log.Logger, apiSecret string) *Bot {
	api := slack.New(apiSecret)
	api.SetDebug(true)

	slack.SetLogger(logger)

	rtm := api.NewRTM()
	go rtm.ManageConnection()

	memory := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	return &Bot{
		logger: logger,
		rtm:    rtm,
		memory: memory,
	}
}
