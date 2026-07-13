package main

import (
	"fmt"
	"time"
)

func fun(s string) {
	for i := 0; i < 3; i++ {
		fmt.Println(s)
		time.Sleep(1 * time.Millisecond)
	}
}

func main() {
	// Direct call
	fun("direct call")

	// TODO: write goroutine with different variants for function call.

	// goroutine function call

	go fun("1-goroutine")

	// goroutine with anonymous function

	go func() {
		fun("2-goroutine")
	}()

	// goroutine with function value call

	fv := fun
	go fv("3-goroutine")

	// wait for goroutines to end

	fmt.Println("wait for go routines")
	time.Sleep(100 * time.Microsecond)

	fmt.Println("done..")
}
