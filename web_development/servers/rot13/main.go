package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

func main() {
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	defer ln.Close()

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatal(err)
		}
		defer conn.Close()

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() { // Read each line of text from the connection. These lines are strings.
		line := strings.ToLower(scanner.Text()) // Set all chars to lower case.
		byteSlice := []byte(line)               // Convert the string to a slice of bytes.
		rot := rot13(byteSlice)

		fmt.Fprintf(conn, "%s - %s\n", line, rot)
	}
}

func rot13(bs []byte) []byte {
	rs := make([]byte, len(bs)) // Initialize a byte slice with an underlying array that is the same length as the argument slice.

	for i, b := range bs {
		if b <= 109 {
			rs[i] = b + 13
		} else {
			rs[i] = b - 13
		}
	}
	return rs
}
