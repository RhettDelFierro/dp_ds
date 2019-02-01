package random

import (
	"fmt"
)

// the SENDER
func f(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

// the RECEIVER
func t() {
	c := make(chan int, 10)
	go f(cap(c), c) // this does not block.

	//THIS will block.
	for i := range c {
		fmt.Println(i)
	}
}

// the SENDER
func f1(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

// the RECEIVER
func t1() {
	c := make(chan int)
	quit := make(chan int)
	// the lambda go function runs on another thread. We have a loop running in the background printing & waiting for
	// f1(..)  - which is initialized at the end of the function.
	go func() { //THIS WILL NOT BLOCK THE CURRENT THREAD. f1(..) blocks the program from exiting.
		// send 10 values to c <-
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		//
		quit <- 0
	}()
	f1(c, quit) // note: this is NOT the goroutine. Remove the "for" logic only leaving the "select" in f1 to see this
}






// My examples - not done.
func f2(n int, c chan int, q chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	q <- 0
}

func t2() {
	c := make(chan int, 10)
	quit := make(chan int)
	go f2(cap(c), c, quit)
	for i := range c {
		fmt.Println(i)
	}
}
