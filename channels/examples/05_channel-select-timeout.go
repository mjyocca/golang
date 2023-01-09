package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	ch1, ch2 := make(chan string), make(chan string)

	timeout := time.After(2 * time.Second)

	go func(ch chan string) {
		for i := 0; i < 2; i++ {
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
			ch <- "message 1"
		}
	}(ch1)

	go func(ch chan string) {
		for i := 0; i < 2; i++ {
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
			ch <- "message 2"
		}
	}(ch2)

	// inifinte loop to read off both channels
	for {
		select {
		case val := <-ch1:
			fmt.Println("Channel 1: ", val)
		case val := <-ch2:
			fmt.Println("Channel 2: ", val)
		case <-timeout:
			fmt.Println("timeout")
			return
		default:
			time.Sleep(time.Millisecond)
		}
	}
}
