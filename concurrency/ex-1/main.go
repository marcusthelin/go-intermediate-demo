package main

import (
	"fmt"
	"sync"
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

	//withWaitGroup()
}

func withWaitGroup() {
	var wg sync.WaitGroup
	wg.Add(2)

	go printMessage("hello", &wg)
	go printMessage("World!!", &wg)

	//go func() {
	//  printMessage("world!!")
	//  wg.Done()
	//}()

	wg.Wait()
	fmt.Println("here...")
}

func getMessages(ch chan string) {
	for i := 1; i < 5; i++ {
		ch <- time.Now().String()
		time.Sleep(time.Millisecond * 500)
	}
	close(ch)
}

func printMessage(msg string, wg *sync.WaitGroup) {
	for i := 1; i <= 5; i++ {
		fmt.Println(i, msg)
		time.Sleep(time.Millisecond * 500)
	}
	wg.Done()
}
