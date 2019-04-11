// map comparation.
package main

import "fmt"

func equal(x, y map[string]int) bool {
    if len(x) != len(y) {
        return false
    }
    for k, xv := range x {
        if yv, ok := y[k]; !ok || xv != yv {
            return false
        }
    }
    return true
}

func main() {
    // True if equal is written incorrectly.
    fmt.Println(equal(map[string]int{"A": 0}, map[string]int{"B": 42}))
}
