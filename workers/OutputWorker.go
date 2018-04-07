package workers

import (
	"log"

	"github.com/nlopes/slack"
)

type OutputWorker struct {
	rtm    *slack.RTM
	logger *log.Logger
}

func (w *OutputWorker) Init(rtm *slack.RTM) {
	w.rtm = rtm
}

func (w *OutputWorker) Process(event slack.RTMEvent) bool {
	w.logger.Print("Event Received: ")

	switch ev := event.Data.(type) {
	case *slack.HelloEvent:
		// Ignore hello

	case *slack.ConnectedEvent:
		w.logger.Println("Infos:", ev.Info)
		w.logger.Println("Connection counter:", ev.ConnectionCount)
		w.rtm.SendMessage(w.rtm.NewOutgoingMessage("Hello world", "general"))

	case *slack.MessageEvent:
		w.logger.Printf("Message: %v\n", ev)

	case *slack.PresenceChangeEvent:
		w.logger.Printf("Presence Change: %v\n", ev)

	case *slack.LatencyReport:
		w.logger.Printf("Current latency: %v\n", ev.Value)

	case *slack.RTMError:
		w.logger.Printf("Error: %s\n", ev.Error())

	case *slack.InvalidAuthEvent:
		w.logger.Printf("Invalid credentials")

	default:
		// Ignore other events..
		w.logger.Printf("Unexpected: %v\n", event.Data)
	}

	return false
}

func (w OutputWorker) Quit() {}

func MakeOutputWorker(logger *log.Logger) Worker {
	return &OutputWorker{
		logger: logger,
	}
}
