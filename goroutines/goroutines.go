package main

import (
	"fmt"
	"math/rand"
	"time"
)

func goRoutine(msg string) {
	for i := 0; i < 10; i++ {
		fmt.Printf("%s - %d\n", msg, i)
		time.Sleep(time.Millisecond)
	}
}

func main() {
	// -----------------------------------
	go goRoutine("async")
	goRoutine("xx")
	goRoutine("yy")

	// -----------------------------------
	messages := make(chan string)
	// fatal error: all goroutines are asleep - deadlock!
	// messages <- "msg"
	go func() {
		time.Sleep(2 * time.Second)
		fmt.Printf("Sending message\n")
		messages <- "msg1"
		messages <- "msg2"
	}()

	fmt.Printf("Waiting for message\n")
	fmt.Printf("Got message: %s\n", <-messages)
	fmt.Printf("Got message: %s\n", <-messages)

	// -----------------------------------
	// make(chan string, 0): Same as make(chan string)
	// make(chan string, 1): Can buffer 1 value without reader
	buffered := make(chan string, 1)
	buffered <- "msg"
	// fatal error: all goroutines are asleep - deadlock!
	//buffered <- "msg"
	msg := <-buffered
	fmt.Printf("Got message: %s\n", msg)

	// -----------------------------------
	approxPi := func(results chan<- float64) {
		// invalid operation: <-results (receive from send-only type chan<- float64)
		// fmt.Printf("Got number: %f\n", <-results)
		var total int64
		var inCircle int64
		for {
			x := rand.Float64()
			y := rand.Float64()
			radiusSqr := x*x + y*y
			if radiusSqr < 1.0 {
				inCircle++
			}
			total++

			if total < 10 || total == 100 || total == 1000 || (total%10000000) == 0 {
				results <- 4.0 * float64(inCircle) / float64(total)
			}
		}
	}
	numbers := make(chan float64, 2)
	go approxPi(numbers)
	for i := 0; i < 20; i++ {
		fmt.Printf("Approx PI: %f\n", <-numbers)
	}
}
