package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

func main() {
	// Listen
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}
	defer li.Close()

	for {
		// Accept (open the connection)
		conn, err := li.Accept()
		if err != nil {
			log.Println(err)
		}
		// Write
		fmt.Fprintln(conn, "The write server accepts a connection and writes back to the caller.")
		io.WriteString(conn, "\nHello. io.WriteString() is one way to write to a net connection.\n")
		fmt.Fprintln(conn, "fmt.Fprintf() is another way to write to the connection.")
		fmt.Fprintf(conn, "%v", "fmt.Fpringtf() is yet another.\n")

		// Close the connection.
		conn.Close()
	}
}
