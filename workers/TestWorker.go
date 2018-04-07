package workers

import (
	"fmt"
	"log"

	"regexp"

	"github.com/nlopes/slack"
)

var testGuidPattern = regexp.MustCompile("test guid")

type TestWorker struct {
	rtm    *slack.RTM
	logger *log.Logger
}

func (w *TestWorker) Init(rtm *slack.RTM) {
	w.rtm = rtm
}

func (w *TestWorker) Process(event slack.RTMEvent) bool {
	switch ev := event.Data.(type) {
	case *slack.MessageEvent:
		fmt.Printf("testing string %s", ev.Text)
		if testGuidPattern.MatchString(ev.Text) {
			fmt.Printf("Message: %v\n", ev)
		}
	}

	return false
}

func (w *TestWorker) Quit() {}

func MakeTestWorker(logger *log.Logger) Worker {
	return &TestWorker{
		logger: logger,
	}
}
