package workers

import "github.com/nlopes/slack"

type Worker interface {
	Init(rtm *slack.RTM)
	Process(event slack.RTMEvent) bool
	Quit()
}
