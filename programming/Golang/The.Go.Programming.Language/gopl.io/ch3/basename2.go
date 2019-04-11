// basename removes directory components and a .suffix.
// e.g., a => a, a.go => a, a/b/c.go => c, a/b.c.go => b.c
package main

import (
    "fmt"
    "strings"
)

func basename(s string) string {
    slash := strings.LastIndex(s, "/")  // -1 if "/" not found
    s = s[slash+1:]
    if dot := strings.LastIndex(s, "."); dot >= 0 {
        s = s[:dot]
    }
    return s
}

func main() {
   fmt.Println(basename("a")) 
   fmt.Println(basename("a.go")) 
   fmt.Println(basename("a/b/c.go")) 
   fmt.Println(basename("a/b.c.go")) 
}
