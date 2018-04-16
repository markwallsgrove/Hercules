package workers

import (
	"log"

	"regexp"

	"github.com/markwallsgrove/hercules/types"
	"github.com/nlopes/slack"
)

// TestWorker demonstration of command execution and approval
type TestWorker struct {
	rtm    types.RTM
	logger *log.Logger
}

// Init the worker
func (w *TestWorker) Init(rtm types.RTM) []Registration {
	w.rtm = rtm

	return []Registration{
		MakeRegistration(
			"hello world",
			regexp.MustCompile("^yo$"),
			w.hello,
		),
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
