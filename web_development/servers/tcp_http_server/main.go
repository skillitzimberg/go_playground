// Create a server that returns the URL of the GET request
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
			continue
		}

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	// Handle requests
	request(conn)

	// Handle response
	response(conn)
}

func request(conn net.Conn) {
	i := 0
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println("Req line: ", ln)
		if i == 0 {

			// Read url from the request line
			mthd := strings.Fields(ln)[0]
			url := strings.Fields(ln)[1]
			vrsn := strings.Fields(ln)[2]
			fmt.Println(mthd, url, vrsn)
		}
		if len(ln) == 0 {
			// The CRFL was found. Headers are done.
			break
		}
		i++
	}

}

func response(conn net.Conn) {
	html := `
		<!DOCTYPE html>
			<html lang="en">
				<head>
					<meta charset="UTF-8">
					<title>HTTP REQUEST: URL</title>
				</head>
				<body>
				<h1>You rang?</h1>
				</body>
			</html>
		`

	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(html))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n") // CRLF (signals the end of response headers)
	fmt.Fprintf(conn, html)
}
