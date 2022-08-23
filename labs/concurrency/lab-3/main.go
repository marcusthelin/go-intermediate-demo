/*
* Given is a producer-consumer scenario, where a producer reads in
* messages from a mockstream and a consumer is processing the
* data. Your task is to change the code so that the producer as well
* as the consumer can run concurrently
 */

package main

import (
	"fmt"
	"time"
)

func producer(stream Stream) (messages []*Message) {
	for {
		message, err := stream.Next()
		if err == ErrEOF {
			return messages
		}

		messages = append(messages, message)
	}
}

func consumer(messages []*Message) {
	for _, t := range messages {
		if t.IsTalkingAboutGo() {
			fmt.Println("\tmessage about golang")
		} else {
			fmt.Println("\tmessage not about golang")
		}
	}
}

func main() {
	start := time.Now()
	stream := GetMockStream()

	// Producer
	messages := producer(stream)

	// Consumer
	consumer(messages)

	fmt.Printf("Process took %s\n", time.Since(start))
}
