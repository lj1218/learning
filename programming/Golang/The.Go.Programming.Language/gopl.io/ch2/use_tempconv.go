package main

import (
    "fmt"
    "gopl.io/ch2/tempconv"
)

func main() {
    fmt.Printf("Brrrr! %v\n", tempconv.AbsoluteZeroC)
    fmt.Println(tempconv.CToF(tempconv.BoilingC))
}
