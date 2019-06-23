// Channels as a handle on a service.
// Our boring function returns a channel that lets us communicate with the boring service it provides.
package main

import (
	"fmt"
	"math/rand"
	"time"
)

// These programs make Joe and Ann count in lockstep.
// We can instead use a fan-in function to let whosever is ready talk. (See demo5)
func main() {
	// We can have more instances of the service.
	joe := boring("Joe")
	ann := boring("Ann")
	for i := 0; i < 5; i++ {
		fmt.Println(<-joe)
		fmt.Println(<-ann)
	}
	fmt.Println("You're boring; I'm leaving.")
}

func boring(msg string) <-chan string { // Returns receive-only channel of strings.
	c := make(chan string)
	go func() { // We launch the goroutine from inside the function.
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d", msg, i) // Expression to be sent can be any suitable value.
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()
	return c // Return the channel to the caller.
}
