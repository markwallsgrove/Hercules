package workers

import (
	"regexp"

	"github.com/nlopes/slack"
)

type Registration struct {
	Name    string
	pattern *regexp.Regexp
	fnc     func(*slack.MessageEvent)
}

func (r *Registration) Apply(event *slack.MessageEvent) bool {
	if r.pattern.MatchString(event.Text) {
		r.fnc(event)
		return true
	}

	return false
}

func MakeRegistration(name string, pattern *regexp.Regexp, fnc func(*slack.MessageEvent)) Registration {
	return Registration{
		Name:    name,
		pattern: pattern,
		fnc:     fnc,
	}
}
