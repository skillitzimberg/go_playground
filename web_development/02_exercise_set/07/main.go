// Building upon the code from the previous problem:
// Print to standard out (the terminal) the REQUEST method and the REQUEST URI from the REQUEST LINE.

// Add this data to your REPONSE so that this data is displayed in the browser.
package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
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

	i := 0
	var rMethod, rURI string

	scanner := bufio.NewScanner(c)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			fmt.Println("Line was an empty string.")
			break
		}
		if i == 0 {
			fields := strings.Fields(line)
			rMethod = fields[0]
			rURI = fields[1]
			fmt.Println("Method:", rMethod, "URI:", rURI)
			i++
		}
		body := "CHECK OUT THE RESPONSE BODY PAYLOAD"
		body += "\n"
		body += rMethod
		body += "\n"
		body += rURI
		io.WriteString(c, "HTTP/1.1 302 Found\r\n")
		fmt.Fprintf(c, "Content-length: %d\r\n", len(body))
		fmt.Fprintf(c, "Content-type: text/plain\r\n")
		io.WriteString(c, "\r\n")

		io.WriteString(c, body)
	}
}
