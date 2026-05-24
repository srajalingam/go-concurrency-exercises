package main

import (
	"io"
	"log"
	"net"
	"os"
)

func main() {
	// connect to server on localhost port 8080
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	// read response from server and print to stdout
	mustCopy(os.Stdout, conn)

}

// mustCopy - utility function
func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
