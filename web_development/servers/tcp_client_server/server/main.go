package server

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func Response() {
	// Listen
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	defer ln.Close()

	// Accept
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Connection opened.")

		// Read & write
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	scanner := bufio.NewScanner(conn)

	for scanner.Scan() {
		line := scanner.Text()
		// Read
		fmt.Printf("The server responds by echoing the client's request: %s\n", line)

		// Write
		fmt.Fprintf(conn, "The server responds by echoing the client's request: %s\n", line)
	}
	defer conn.Close()

	fmt.Println("Connection closed.")
}
