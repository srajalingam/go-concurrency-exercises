package main

import (
	"bytes"
	"io"
	"os"
	"sync"
	"time"
)

// create pool of bytes.Buffers which can be reused.
var bufferPool = sync.Pool{
	New: func() interface{} {
		println("Creating new buffer")
		return new(bytes.Buffer)
	},
}

func log(w io.Writer, val string) {
	var b = bufferPool.Get().(*bytes.Buffer)

	b.Reset()

	b.WriteString(time.Now().Format("15:04:05"))
	b.WriteString(" : ")
	b.WriteString(val)
	b.WriteString("\n")

	w.Write(b.Bytes())

	bufferPool.Put(b)
}

func main() {
	log(os.Stdout, "debug-string1")
	log(os.Stdout, "debug-string2")
}
