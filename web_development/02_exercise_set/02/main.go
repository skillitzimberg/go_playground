// EXERCISE 02
// Building upon the code from the previous exercise:
// In that previous exercise, we WROTE to the connection.

// Now I want you to READ from the connection.

// You can READ and WRITE to a net.Conn as a connection implements both the reader and writer interface.

// Use bufio.NewScanner() to read from the connection.

// After all of the reading, include these lines of code:

// fmt.Println("Code got here.") io.WriteString(c, "I see you connected.")

// Launch your TCP server.

// In your web browser, visit localhost:8080.

// Now go back and look at your terminal.

// Can you answer the question as to why "I see you connected." is never written?
package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
)

func main() {
	lstnr, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	defer lstnr.Close()

	for {
		conn, err := lstnr.Accept()
		if err != nil {
			log.Fatal(err)
		}
		defer conn.Close()

		scanner := bufio.NewScanner(conn)
		for scanner.Scan() {
			line := scanner.Text()
			fmt.Println(line)
		}
		// This scanner/reader never stops. It won't stop unless it receives an error or io.EOF (an end of file).

		// That's why this code doesn't get run until the connection is broken.
		fmt.Println("Code got here.")

		// And with that connection broken, there is no connection to write back to here. ???
		io.WriteString(conn, "I see you connected.")
	}
}
