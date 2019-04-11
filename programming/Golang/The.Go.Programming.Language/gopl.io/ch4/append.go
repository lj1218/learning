package main

import "fmt"

func appendInt(x []int, y ...int) []int {
    var z []int
    zlen := len(x) + len(y)
    if zlen <= cap(x) {
        // There is room to grow. Extend the slice.
        z = x[:zlen]
    } else {
        // There is insufficient space. Allocate a new array.
        // Grow by doubling, for amortized linear complexity.
        zcap := zlen
        if zcap < 2*len(x) {
            zcap = 2 * len(x)
        }
        z = make([]int, zlen, zcap)
        copy(z, x)  // a built-in function; see text
    }
    //z[len(x)] = y
    copy(z[len(x):], y)
    return z
}

func main() {
    var x, y []int
    for i := 0; i < 10; i++ {
        if i < 9 {
            y = appendInt(x, i)
        } else {
            y = appendInt(x, i, i+1)
        }
        fmt.Printf("%d  cap=%d\t%v\n", i, cap(y), y)
        x = y
    }
}
