package random

import "fmt"

// The select statement lets a goroutine wait on multiple communication operations.
func fibonacci1(c, quit chan int) { //NOTE THIS IS BOHT A SENDER AND A RECEIVER
	x, y := 0, 1
	for {
		// A select blocks UNTIL one of its cases can run, then it executes that case.
		// It chooses one at RANDOM if multiple are ready.
		select {
		case c <- x: //send the value to the channel       //A RECEIVER
			x, y = y, x+y // reassign values.
		case <-quit:                                       //A SENDER
			fmt.Println("quit")
			return
		}
	}
}

// <-c is the RECEIVER.
func test1() {
	c := make(chan int)
	quit := make(chan int)
	// the goroutine runs ONCE in the background - it's a single function call - but does not block the program from terminating.
	go func() { //NOTE: THIS IS A SENDER AND A RECEIVER
		for i := 0; i < 10; i++ { //will make 10 fmt.Println's RECEIVING a (<-c) SIGNAL. This logic says that we're going to RECEIVE 10 signals, then exit the loop and wait for a quit.
			fmt.Println(<-c) //RECEIVE
		}
		quit <- 0 //SEND
	}()
	// initiate the SENDER
	fibonacci1(c, quit)
}
