package main

import (
    "fmt"
    "io/ioutil"
    "os"
    "strings"
)

func main() {
    counts := make(map[string]int)
    for _, filename := range os.Args[1:] {
        data, err := ioutil.ReadFile(filename)
        if err != nil {
            fmt.Fprintf(os.Stderr, "dup3: %v\n", err)
            continue
        }
//        if data[len(data) - 1] == '\n' {
//            data[len(data) - 1] = 0
//        }
        for _, line := range strings.Split(string(data), "\n") {
            if len(line) > 1 {
                counts[line]++
            }
        }
    }

    for line, count := range counts {
        fmt.Printf("%d\t%s\n", count, line)
    }
}
