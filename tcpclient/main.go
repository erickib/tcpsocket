package main

import (
	"fmt"
	"log"
	"net"
)

func main() {

	// create a tcp socket connection on localhost at port 1000
	conn, err := net.Dial("tcp", "127.0.0.1:1000")
	if err != nil {
		log.Fatal(err)
	}

	var b = []byte("Hello, server!")
	// write some data to the socket
	// returns the number of bytes written from b (0 <= n <= len(b))
	n, err := conn.Write(b)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Number of bytes written: %d", n)

	// Close the connection
	conn.Close()
}
