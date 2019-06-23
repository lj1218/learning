// Quit channel.
// We can turn this around and tell Joe to stop when we're tired of listening to him.
// How do we know it's finished? Wait for it to tell us it's done: receive on the quit channel.
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	quit := make(chan string)
	c := boring("Joe", quit)
	for i := rand.Intn(10); i >= 0; i-- {
		fmt.Println(<-c)
	}
	quit <- "Bye!"
	fmt.Printf("Joe says: %q\n", <-quit)
}

func boring(msg string, quit chan string) <-chan string { // Returns receive-only channel of strings.
	c := make(chan string)
	go func() { // We launch the goroutine from inside the function.
		for i := 0; ; i++ {
			select {
			case c <- fmt.Sprintf("%s %d", msg, i): // Expression to be sent can be any suitable value.
				// do nothing
			case <-quit:
				cleanup()
				quit <- "See you!"
				return
			}
			time.Sleep(time.Duration(rand.Intn(2e3)) * time.Millisecond)
		}
	}()
	return c // Return the channel to the caller.
}

func cleanup() {
	fmt.Println("cleanup...")
}
