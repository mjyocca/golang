package main

import "fmt"

func main() {

	messages := make(chan string)

	// if we sent to the channel in same routine, would block and end up causing a deadlock
	go func() { messages <- "ping" }()

	msg := <-messages
	fmt.Println(msg)
}
