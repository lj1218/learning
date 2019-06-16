// 9.8.3. GOMAXPROCS
package main

import "fmt"

func main() {
	for {
		go fmt.Print(0)
		fmt.Print(1)
	}
}
// GOMAXPROCS=1 go run gomaxprocs.go
// GOMAXPROCS=2 go run gomaxprocs.go
