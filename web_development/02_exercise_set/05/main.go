// Building upon the code from the previous problem:
// Add code to WRITE to the connection.
package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
)

func main() {
	lsntr, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	defer lsntr.Close()

	for {
		conn, err := lsntr.Accept()
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
			fmt.Println("The line was an empty string.")
			break
		}
		io.WriteString(c, line)
	}

	fmt.Println("Code got here.")
	io.WriteString(c, "The connection will be closed now.")
}
