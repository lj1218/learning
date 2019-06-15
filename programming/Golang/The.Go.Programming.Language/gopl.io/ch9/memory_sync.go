package main

import (
	"fmt"
	"sync"
)

func main() {
	var x, y int
	var wg sync.WaitGroup

	wg.Add(2)
	go func() {
		defer wg.Done()
		x = 1                   // A1
		fmt.Print("y:", y, " ") // A2
	}()

	go func() {
		defer wg.Done()
		y = 1                   // B1
		fmt.Print("x:", x, " ") // B2
	}()

	wg.Wait()
	fmt.Println()
}
