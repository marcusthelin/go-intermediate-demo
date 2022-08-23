/**
* DO NOT EDIT THIS PART
* Your task is to edit `main.go`
 */

package main

import (
	"errors"
	"strings"
	"time"
)

// GetMockStream is a blackbox function which returns a mock stream for
// demonstration purposes
func GetMockStream() Stream {
	return Stream{0, mockdata}
}

// Stream is a mock stream for demonstration purposes, not threadsafe
type Stream struct {
	pos      int
	messages []Message
}

// ErrEOF returns on End of File error
var ErrEOF = errors.New("End of File")

// Next returns the next Message in the stream, returns EOF error if
// there are no more messages
func (s *Stream) Next() (*Message, error) {

	// simulate delay
	time.Sleep(320 * time.Millisecond)
	if s.pos >= len(s.messages) {
		return &Message{}, ErrEOF
	}

	message := s.messages[s.pos]
	s.pos++

	return &message, nil
}

// Message defines the simlified representation of a message
type Message struct {
	Text string
}

// IsTalkingAboutGo is a mock process which pretend to be a sophisticated procedure to analyse whether message is talking about go or not
func (t *Message) IsTalkingAboutGo() bool {
	// simulate delay
	time.Sleep(330 * time.Millisecond)

	hasGolang := strings.Contains(strings.ToLower(t.Text), "golang")
	hasGopher := strings.Contains(strings.ToLower(t.Text), "gopher")

	return hasGolang || hasGopher
}

var mockdata = []Message{
	{
		"#golang top tip: if your unit tests import any other package you wrote, including themselves, they're not unit tests.",
	}, {
		"Backend developer, doing frontend featuring the eternal struggle of centering something. #coding",
	}, {
		"Re: Popularity of Golang in China: My thinking nowadays is that it had a lot to do with this book and author https://github.com/astaxie/build-web-application-with-golang",
	}, {
		"Looking forward to the #gopher meetup in Hsinchu tonight with @ironzeb!",
	}, {
		"I just wrote a golang slack bot! It reports the state of github repository. #Slack #golang",
	},
}
