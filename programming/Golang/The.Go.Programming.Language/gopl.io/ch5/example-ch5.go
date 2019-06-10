package main

import "fmt"

func main() {
    // 5.6.1 Caveat: Capturing Iteration Variables
    example_5_6_1_1()
    example_5_6_1_2()
}

func example_5_6_1_1() {
    var printVals []func()
    vals := []string{"a", "b", "c", "d", "e"}
    for _, v := range vals {
        val := v  // NOTE: necessary!
        printVals = append(printVals, func() {
            fmt.Print(val, " ")
        })
    }

    for _, printVal := range printVals {
        printVal()  // print value 
    }

    fmt.Println()
}

func example_5_6_1_2() {
    var printVals []func()
    vals := []string{"a", "b", "c", "d", "e"}
    for _, val := range vals {
        printVals = append(printVals, func() {
            fmt.Print(val, " ")  // NOTE: incorrect!
        })
    }

    for _, printVal := range printVals {
        printVal()  // print value 
    }

    fmt.Println()
}

