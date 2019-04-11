package main

import (
    "fmt"
    "sort"
)

func main() {
    // 4.1. Arrays
    var a [3]int              // arrays of 3 integers
    fmt.Println(a[0])         // print the first element
    fmt.Println(a[len(a)-1])  // print the last element, a[2]
    // Print the indices and elements.
    for i, v := range a {
        fmt.Printf("%d %d\n", i, v)
    }
    // Print the elements only.
    for _, v := range a {
        fmt.Printf("%d\n", v)
    }

    var q [3]int = [3]int{1, 2, 3}
    var r [3]int = [3]int{1, 2}
    fmt.Println(q[2])  // "3"
    fmt.Println(r[2])  // "0"
    q2 := [...]int{1, 2, 3}
    fmt.Printf("%T\n", q2)  // "[3]int"
    //q = [4]int{1, 2, 3, 4}  // cannot use [4]int literal (type [4]int) as type [3]int in assignment

    type Currency int
    const (
        USD Currency = iota
        EUR
        GBP
        RMB
    )
    symbol := [...]string{USD: "$", EUR: "%", GBP: "^", RMB: "￥"}
    fmt.Println(RMB, symbol[RMB])  // "3 ￥"
    r2 := [...]int{99: -1}
    fmt.Printf("%d %d\n", r2[98], r2[99])  // "0 -1"

    // Arrays comparation
    a_ := [2]int{1, 2}
    b := [...]int{1, 2}
    c := [2]int{1, 3}
    fmt.Println(a_ == b, a_ == c, b == c)  // "true false false"
    //d := [3]int{1, 2}
    //fmt.Println(a_ == d) // compile error: invalid operation: a_ == d (mismatched types [2]int and [3]int)


    // 4.2. Slices
    months := [...]string{1: "January", 2: "February", 3: "March", 4: "April",
        5: "May", 6: "June", 7: "July", 8: "August", 9: "September", 10: "October",
        11: "November", 12: "December"}
    Q2 := months[4:7]
    summer := months[6:9]
    fmt.Println(Q2)      // [April May June]
    fmt.Println(summer)  // [June July August]
    for _, s := range summer {
        for _, q := range Q2 {
            if s == q {
                fmt.Printf("%s appears in both\n", s)
            }
        }
    }
    // extend a slice (within capacity)
    //fmt.Println(summer[:20])  // panic: runtime error: slice bounds out of range
    endlessSummer := summer[:5]
    fmt.Println(endlessSummer)  // [June July August September October]

    // 4.2.1. The append Function
    var runes []rune
    for _, r := range "Hello, 世界" {
        runes = append(runes, r)
    }
    fmt.Printf("%q\n", runes)
    var x []int
    x = append(x, 1)
    x = append(x, 2, 3)
    x = append(x, 4, 5, 6)
    x = append(x, x...)  // append the slice x
    fmt.Println(x)       // [1 2 3 4 5 6 1 2 3 4 5 6]


    // 4.3. Maps
    // ages := make(map[string]int)  // mapping from strings to ints
    // ages["alice"] = 31
    // ages["charlie"] = 34
    ages := map[string]int{
        "alice":   31,
        "charlie": 34,
    }
    fmt.Println(ages)  // map[alice:31 charlie:34]
    fmt.Println(ages["alice"])  // 31
    delete(ages, "alice")
    fmt.Println(ages["alice"])  // 0
    ages["bob"] += 1
    fmt.Println(ages)  // map[bob:1 charlie:34]
    ages["bob"]++
    fmt.Println(ages)  // map[bob:2 charlie:34]
    // _ = &ages["bob"]  // compile error: cannot take the address of map element
    ages["alice"] = 31
    ages["lily"] = 19
    for name, age := range ages {
        fmt.Printf("%s\t%d\n", name, age)
    }
    fmt.Println("-----")

    var names[] string
    for name := range ages {
        names = append(names, name)
    }
    sort.Strings(names)
    for _, name := range names {
        fmt.Printf("%s\t%d\n", name, ages[name])
    }
 
    //age, ok := ages["Joe"]
    if age, ok := ages["Joe"]; !ok {
        fmt.Printf("Joe is not a key in map ages, and age = %d\n", age)
    }
}
