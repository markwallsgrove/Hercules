package workers

import (
	"log"
	"regexp"

	"github.com/nlopes/slack"
)

type Registration struct {
	name    string
	pattern *regexp.Regexp
	fnc     func(*slack.MessageEvent)
}

func (r *Registration) Apply(event *slack.MessageEvent) bool {
	log.Println("checking message ", event.Text)
	if r.pattern.MatchString(event.Text) {
		log.Println("message matches")
		r.fnc(event)
		return true
	}

	return false
}
