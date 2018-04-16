package bot

// func TestBlah(t *testing.T) {
// 	// how long should we wait for this test to finish?
// 	// maxWait := 5 * time.Second
// 	// start our test server
// 	s := slacktest.NewTestServer()
// 	go s.Start()

// 	bot := StartBot("secret", s.GetAPIURL())

// 	// create a channel to pass our results from the next goroutine
// 	// that is a goroutine doing the normal range over rtm.IncomingEvents
// 	// messageChan := make(chan (*slack.MessageEvent), 1)

// 	// since we want to test direct messages, let's send one to the bot
// 	s.SendDirectMessageToBot(t.Name())

// 	// // now we block this test
// 	// select {
// 	// // if we get a slack.MessageEvent, perform some assertions
// 	// case _ = <-messageChan:
// 	// 	// assert.Equal(t, "D024BE91L", m.Channel)
// 	// 	// assert.Equal(t, t.Name(), m.Text)
// 	// 	break
// 	// // if we hit our timeout, fail the test
// 	// case <-time.After(maxWait):
// 	// 	// assert.FailNow(t, "did not get direct message in time")
// 	// 	t.FailNow()
// 	// }

// 	bot.Quit()
// }
