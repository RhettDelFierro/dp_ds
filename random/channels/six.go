package random

import (
	"fmt"
	"time"
)

func main() {
	tick := time.Tick(100 * time.Millisecond) // returns <-chan
	boom := time.After(500 * time.Millisecond) // returns <-chan
	for {
		select {
		case <-tick:
			fmt.Println("tick.")
		case <-boom:
			fmt.Println("BOOM!")
			return //here is the return to terminate the function.
		default:
			fmt.Println("    .")
			time.Sleep(50 * time.Millisecond)
		}
	}
}
