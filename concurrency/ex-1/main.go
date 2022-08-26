package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)
	go getMessages(ch)

	for msg := range ch {
		//msg, open := <-ch
		//if !open {
		//  break
		//}
		fmt.Println(msg)
	}

}

func getMessages(ch chan string) {
	for i := 1; i < 5; i++ {
		ch <- time.Now().String()
		time.Sleep(time.Millisecond * 500)
	}
	close(ch)
}
