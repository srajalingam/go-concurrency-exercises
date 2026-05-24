package main

import (
	"fmt"
	"io"
	"net"
	"time"
)

func main() {
	//write server program to handle concurrent client connections.
	ln, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		panic(err)
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			panic(err)
		}
		fmt.Println("Client connected:", conn.RemoteAddr())
		go handleConn(conn)
	}
}

// handleConn - utility function
func handleConn(c net.Conn) {
	defer c.Close()
	fmt.Println("server connected")
	for {
		_, err := io.WriteString(c, "response from server\n")
		if err != nil {
			return
		}
		time.Sleep(time.Second * 2)
	}
}
