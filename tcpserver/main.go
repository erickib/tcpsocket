package main

import (
	"fmt"
	"log"
	"net"
)

func main() {

	//listen for incoming connections on localhost at port 1000
	ln, err := net.Listen("tcp", "127.0.0.1:1000")
	if err != nil {
		log.Fatal(err)
	}

	//loop through accepting incoming connections and handle them
	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}

		//handle connection spawning a new goroutine
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {

	//defer will be called in the end of this function scope
	//and close this connection
	defer conn.Close()

	//read data from socket to buffer
	buf := make([]byte, 1024)
	_, err := conn.Read(buf)
	if err != nil {
		log.Fatal(err)
	}

	//print data
	fmt.Printf("received: %s", buf)
}
