package workers

import (
	"log"
	"os"
	"testing"

	"github.com/nlopes/slack"
)

type MockRTM struct {
	Outgoing []*slack.OutgoingMessage
}

func (r *MockRTM) SendMessage(msg *slack.OutgoingMessage) {
	r.Outgoing = append(r.Outgoing, msg)
}

func (r *MockRTM) NewOutgoingMessage(text string, channelID string) *slack.OutgoingMessage {
	return &slack.OutgoingMessage{
		ID:      0,
		Channel: channelID,
		Text:    text,
	}
}

func (r MockRTM) Disconnect() error {
	return nil
}

func TestTestWorker(t *testing.T) {
	rtm := &MockRTM{[]*slack.OutgoingMessage{}}
	logger := log.New(os.Stdout, "bot: ", log.Lshortfile|log.LstdFlags)
	testWorker := MakeTestWorker(logger)
	listeners := testWorker.Init(rtm)

	if len(listeners) != 1 {
		t.Error("unexpected amount of listeners")
		t.FailNow()
	}

	event := &slack.MessageEvent{
		Msg:        slack.Msg{Text: "yo", Channel: "ops"},
		SubMessage: nil,
	}

	listener := listeners[0]
	if listener.Apply(event) == false {
		t.Error("message was not processed")
		t.FailNow()
	}

	if len(rtm.Outgoing) != 1 {
		t.Error("Unexpected outgoing length")
		t.FailNow()
	}

	outgoing := rtm.Outgoing[0]
	if outgoing.Text != "hey!" || outgoing.Channel != "ops" {
		t.Errorf("Unexpected message: %v", outgoing)
	}
}
