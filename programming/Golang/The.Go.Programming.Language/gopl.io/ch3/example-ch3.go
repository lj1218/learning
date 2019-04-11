package main

import (
    "fmt"
    "strconv"
    "unicode/utf8"
)

func main() {
    var u uint8 = 255
    fmt.Println(u, u+1, u*u)  // "255 0 1"

    var i int8 = 127
    fmt.Println(i, i+1, i*i)  // "127 -128 1"

    var x uint8 = 1<<1 | 1<<5  // 00100010
    var y uint8 = 1<<1 | 1<<2  // 00000110
    fmt.Printf("%08b\n", x)
    fmt.Printf("%08b\n", y)

    fmt.Printf("%08b\n", x&y)  // 00000010
    fmt.Printf("%08b\n", x|y)  // 00100110
    fmt.Printf("%08b\n", x^y)  // 00100100
    fmt.Printf("%08b\n", x&^y) // 00100000

    for i := uint(0); i < 8; i++ {
        if x&(1<<i) != 0 {
            fmt.Println(i) // 1,5
        }
    }
    fmt.Printf("%08b\n", x<<1)  // 01000100
    fmt.Printf("%08b\n", x>>1)  // 00010001

    o := 0666
    fmt.Printf("%d %[1]o %#[1]o\n", o)  // "438 666 0666"
    x2 := int64(0xdeadbeef)
    fmt.Printf("%d %[1]x %#[1]x %#[1]X\n", x2) // 3735928559 deadbeef 0xdeadbeef 0XDEADBEEF

    ascii := 'a'
    unicode := '国'
    newline := '\n'
    fmt.Printf("%d %[1]c %[1]q\n", ascii)  // 97 a 'a'
    fmt.Printf("%d %[1]c %[1]q\n", unicode)  // 22269 国 '国'
    fmt.Printf("%d %[1]q\n", newline)  // 10 '\n'

    s := "Hello, 世界"
    fmt.Println(len(s))  // 13
    fmt.Println(utf8.RuneCountInString(s))  // 9
    for i := 0; i < len(s); {
        r, size := utf8.DecodeRuneInString(s[i:])
        fmt.Printf("%d\t%c\n", i, r)
        i += size
    }
    for i, r := range s {
        fmt.Printf("%d\t%q\t%d\n", i, r, r)
    }

    diamond_like := '\uFFFD'
    fmt.Printf("%c\n", diamond_like)

//    s = "世界"
    fmt.Printf("% x\n", s)
    r := []rune(s)
    fmt.Printf("%x\n", r)
    fmt.Println(string(r))
    fmt.Println(string(65))
    fmt.Println(string(0x4eac))
    fmt.Println(string(1234567))

    x3 := 123
    y3 := fmt.Sprintf("%d", x3)
    fmt.Println(y3, strconv.Itoa(x3))  // "123 123"

    fmt.Println(strconv.FormatInt(int64(x3), 2))  // "1111011"
    fmt.Println(fmt.Sprintf("x3=%b", x3))  // "x3=1111011"
    x4, err4 := strconv.Atoi("123")  // x is an int
    y5, err5 := strconv.ParseInt("123", 10, 64)  // base 10, upto 64 bits
    fmt.Println(x4, err4)
    fmt.Println(y5, err5)

    // 3.6. Constants
    const (
        a = 1
        b
        c = 2
        d
        e = "hhh"
        f
        g = 3
        h
    )
    fmt.Println(a, b, c, d, e, f, g, h)  // "1 1 2 2 hhh hhh 3 3"

    // 3.6.1. The Constant Generator iota
    type Weekday int
    const (
        Sunday Weekday = iota
        Monday
        Tuesday
        Wednesday
        Thursday
        Friday
        Saturday
    )
    fmt.Println(Sunday, Monday, Tuesday, Wednesday, Thursday, Friday, Saturday)  // "0 1 2 3 4 5 6"

    const (
        _ = 1 << (10 * iota)
        KiB // 1 << 10 = 1024
        MiB // 1 << 20 = 1048576
        GiB // 1 << 30 = 1073741824
        TiB // 1 << 40 = 1099511627776 (exceeds 1 << 32)
        PiB // 1 << 50 = 1125899906842624
        EiB // 1 << 60 = 1152921504606846976
        ZiB // 1 << 70 = 1180591620717411303424 (exceeds 1 << 64)
        YiB // 1 << 80 = 1208925819614629174706176
    )
    fmt.Println(KiB, MiB, GiB, TiB, PiB, EiB)
    fmt.Println(YiB / ZiB)  // "1024"

    // 3.6.2. Untyped Constants
    fmt.Printf("%T\n", 0)       // "int"
    fmt.Printf("%T\n", 0.0)     // "float64"
    fmt.Printf("%T\n", 0i)      // "complex128"
    fmt.Printf("%T\n", '\000')  // "int32" (rune)
}
