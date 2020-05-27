// Building upon the code from the previous problem:
// Change your RESPONSE HEADER "content-type" from "text/plain" to "text/html"

// Change the RESPONSE from "CHECK OUT THE RESPONSE BODY PAYLOAD" (and everything else it contained: request method, request URI) to an HTML PAGE that prints "HOLY COW THIS IS LOW LEVEL" in tags.
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
		}

		go serve(conn)
	}
}

func serve(c net.Conn) {
	defer c.Close()

	rLine := 0
	var rMethod, rURI string

	scanner := bufio.NewScanner(c)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			fmt.Println("Line was an empty string.")
			break
		}

		if rLine == 0 {
			fields := strings.Fields(line)
			rMethod = fields[0]
			rURI = fields[1]
		}

		html := "<h1>The request method was %s. The Requested resource was %s</h1>"
		body := fmt.Sprintf(html, rMethod, rURI)
		io.WriteString(c, "GET http/1.1 302 found\r\n")
		fmt.Fprintf(c, "Content-length: %d\r\n", len(body))
		fmt.Fprintf(c, "Content-type: text/html\r\n")
		io.WriteString(c, "\r\n")

		io.WriteString(c, body)
	}

}
