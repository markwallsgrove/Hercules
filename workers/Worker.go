package workers

import "github.com/markwallsgrove/hercules/types"

type Worker interface {
	Init(rtm types.RTM) []Registration
	Quit()
}
