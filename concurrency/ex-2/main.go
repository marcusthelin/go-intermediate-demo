package main

import "fmt"

// channel and buffered channel example
func main() {
	ch := make(chan string, 3)

	ch <- "hello"
	ch <- "world!!"

	msg := <-ch
	fmt.Print(msg)

	msg = <-ch
	fmt.Print(msg)

	//msg = <-ch
	//fmt.Print(msg)
}
