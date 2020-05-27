// Building upon the code from the previous problem:
// Extract the code you wrote to READ from the connection using bufio.NewScanner into its own function called "serve".

// Pass the connection of type net.Conn as an argument into this function.

// Add "go" in front of the call to "serve" to enable concurrency and multiple connections.

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
			continue
		}

		go serve(conn)
	}
}

func serve(c net.Conn) {
	defer c.Close()

	scanner := bufio.NewScanner(c)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			fmt.Println("Line was an empty string.")
			break
		}

		fmt.Println(line)
	}

	fmt.Println("Code got here.")
	io.WriteString(c, "Connection acknowledged and terminated.")
}
