package main

import (
	"bytes"
	"fmt"
	"io"
)

const debug = false

func main() {
	var buf *bytes.Buffer
	//var buf io.Writer
	fmt.Printf("buf: %v, %T, buf == nil: %v\n", buf, buf, buf == nil)
	if debug {
		buf = new(bytes.Buffer)  // 启用输出收集
	}
	fmt.Printf("buf: %v, %T, buf == nil: %v\n", buf, buf, buf == nil)
	f(buf)
	if debug {
		// ... 使用 buf ...
	}
}

func f(out io.Writer) {
	// ... 其他代码 ...
	fmt.Printf("out: %v, %T, out == nil: %v\n", out, out, out == nil)
	if out != nil {
		_, _ = out.Write([]byte("done!\n"))
	}
}
