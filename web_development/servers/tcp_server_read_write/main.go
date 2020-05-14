package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {
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
		// Read
		line := scanner.Text()
		fmt.Println(line) // This lets the user see what they are typing into the terminal.

		// Write
		fmt.Fprintf(conn, "Echo: %s\n", line)
	}
	defer conn.Close()

	// we never get here
	// we have an open stream connection
	// how does the above reader know when it's done?
	fmt.Println("Connection closed.")
}
