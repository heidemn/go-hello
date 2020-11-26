package main

import (
	"fmt"
	"time"
)

func goRoutine(msg string) {
	for i := 0; i < 10; i++ {
		fmt.Printf("%s - %d\n", msg, i)
		time.Sleep(time.Millisecond)
	}
}

func main() {
	go goRoutine("async")
	goRoutine("xx")
	goRoutine("yy")

	messages := make(chan string)
	// fatal error: all goroutines are asleep - deadlock!
	// messages <- "msg"
	go func() {
		time.Sleep(2 * time.Second)
		fmt.Printf("Sending message\n")
		messages <- "msg"
	}()

	fmt.Printf("Waiting for message\n")
	msg := <-messages
	fmt.Printf("Got message: %s\n", msg)
}
