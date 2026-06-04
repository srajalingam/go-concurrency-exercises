package main

import (
	"time"
)

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(3 * time.Second)
		ch1 <- "one"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		ch2 <- "two"
	}()

	// Use select to wait for messages from both channels
	//DESCRIPTION: Use select to wait for messages from both channels and print them as they arrive.
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-ch1:
			println("Received from ch1:", msg1)
		case msg2 := <-ch2:
			println("Received from ch2:", msg2)
		}
	}

	// msg1 := <-ch1
	// println("Received from ch1:", msg1)

	// msg2 := <-ch2
	// println("Received from ch2:", msg2)
}
