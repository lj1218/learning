// Exercise 8.8:
//   Using a select statement, add a timeout to the echo server from Section 8.3 so
// that it disconnects any client that shouts not hing wit hin 10 seconds.
package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"sync"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err) // e.g., connection aborted
			continue
		}
		go handleConn(conn) // handle connections concurrently
	}
}

func echo(c net.Conn, shout string, delay time.Duration) {
	if len(shout) == 0 {
		return
	}
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
}

func handleConn(c net.Conn) {
	input := bufio.NewScanner(c)

	msg := make(chan string)
	disconn := make(chan struct{})

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		ticker := time.NewTicker(10 * time.Second)
		defer wg.Done()
		defer ticker.Stop()
		echoCalled := false
		for {
			select {
			case <-ticker.C:
				if !echoCalled {
					fmt.Println("timed out")
					return
				}
				echoCalled = false
			case shout := <-msg:
				echoCalled = true
				wg.Add(1)
				go func() {
					defer wg.Done()
					echo(c, shout, 1*time.Second)
				}()
			case <-disconn:
				fmt.Println("disconnected by client")
				return
			}
		}
	}()

	go func() {
		for input.Scan() {
			msg <- input.Text()
		}
		disconn <- struct{}{}
	}()

	wg.Wait()
	// NOTE: ignoring potential errors from input.Err()
	c.Close()
}
