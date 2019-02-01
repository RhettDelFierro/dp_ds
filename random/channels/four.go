package random

import (
	"fmt"
)

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	// Only the sender (ch <- val) should close a channel, never the receiver.
	// Sending on a closed channel will cause a panic: try putting close(c) above to loop to see this.
	close(c) //If you delete this, the test() function "range" will never terminate (deadlock)
				// fatal error: all goroutines are asleep - deadlock!
				// which makes sense because the goroutine is finished, but range is never told to close.
				// close(..) is specifically for "range"s.
}

func test() {
	c := make(chan int, 10)
	go fibonacci(cap(c), c)
	// The loop for i := range c RECEIVES values from the channel repeatedly UNTIL IT IS CLOSED.
	for i := range c {
		fmt.Println(i)
	}
}

/*
* Channels aren't like files; you don't usually need to close them.
* Closing is only necessary when the RECEIVER (val <- ch || for i := range c) must be TOLD
* there are no more values coming, such as to terminate a RANGE loop.
*/