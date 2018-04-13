package workers

import (
	"log"

	"regexp"

	"github.com/go-redis/redis"
	"github.com/nlopes/slack"
)

// TestWorker demonstration of command execution and approval
type TestWorker struct {
	rtm    *slack.RTM
	logger *log.Logger
	memory *redis.Client
}

// Init the worker
func (w *TestWorker) Init(rtm *slack.RTM, memory *redis.Client) []Registration {
	w.rtm = rtm
	w.memory = memory

	return []Registration{
		Registration{"hello world", regexp.MustCompile("^yo$"), w.hello},
	}
}

func (w *TestWorker) hello(event *slack.MessageEvent) {
	w.rtm.SendMessage(w.rtm.NewOutgoingMessage("hey!", event.Channel))
}

// Quit the worker and close down any resources
func (w *TestWorker) Quit() {}

// MakeTestWorker constructor
func MakeTestWorker(logger *log.Logger) Worker {
	return &TestWorker{
		logger: logger,
	}
}
