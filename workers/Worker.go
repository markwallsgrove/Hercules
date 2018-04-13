package workers

import (
	"github.com/go-redis/redis"
	"github.com/nlopes/slack"
)

type Worker interface {
	Init(rtm *slack.RTM, redis *redis.Client) []Registration
	Quit()
}
