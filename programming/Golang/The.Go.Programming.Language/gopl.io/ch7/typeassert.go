package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

func main() {
	var w io.Writer
	w = os.Stdout
	_ = w.(*os.File)
	//_ = w.(*bytes.Buffer) // panic: interface conversion: io.Writer is *os.File, not *bytes.Buffer
	if _, ok := w.(*bytes.Buffer); !ok {
		fmt.Println("error: w.(*bytes.Buffer)")
	}

	rw := w.(io.ReadWriter)
	fmt.Printf("rw: %T, %v\n", rw, rw)

	_, err := os.Open("/no/such/file")
	fmt.Println(err)         // open /no/such/file: no such file or directory
	fmt.Printf("%#v\n", err) // &os.PathError{Op:"open", Path:"/no/such/file", Err:0x2}
}
