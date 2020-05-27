// Building upon the code from the previous problem:
// Add code to respond to the following METHODS & ROUTES: GET / GET /apply POST /apply
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

	scanner := bufio.NewScanner(c)
	rLine := 0
	var rMethod, rURI, req string

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			fmt.Println("Line was an empty string.")
			break
		}

		if rLine == 0 {
			rLineFields := strings.Fields(line)
			rMethod = rLineFields[0]
			rURI = rLineFields[1]
			req = fmt.Sprintf("%s %s", rMethod, rURI)
			rLine++
		}

		switch req {
		case "GET /":
			handleIndex(c)
		case "GET /apply":
			handleApply(c)
		case "POST /apply":
			handleApplyPost(c)
		default:
			handleDefault(c)
		}
	}
}

func handleIndex(c net.Conn) {
	body := `
		<!DOCTYPE html>
		<html lang="en">
		<head>
			<meta charset="UTF-8">
			<title>GET INDEX</title>
		</head>
		<body>
			<h1>"GET INDEX"</h1>
			<a href="/">index</a><br>
			<a href="/apply">apply</a><br>
		</body>
		</html>
	`
	handleResponseHeaders(c, body)
}

func handleApply(c net.Conn) {
	body := `
		<!DOCTYPE html>
		<html lang="en">
		<head>
			<meta charset="UTF-8">
			<title>GET DOG</title>
		</head>
		<body>
			<h1>"GET APPLY"</h1>
			<a href="/">index</a><br>
			<a href="/apply">apply</a><br>
			<form action="/apply" method="POST">
			<input type="hidden" value="In my good death">
			<input type="submit" value="submit">
			</form>
		</body>
		</html>
	`
	handleResponseHeaders(c, body)
}

func handleApplyPost(c net.Conn) {
	body := `
		<!DOCTYPE html>
		<html lang="en">
		<head>
			<meta charset="UTF-8">
			<title>POST APPLY</title>
		</head>
		<body>
			<h1>"POST APPLY"</h1>
			<a href="/">index</a><br>
			<a href="/apply">apply</a><br>
		</body>
	</html>
	`
	handleResponseHeaders(c, body)
}

func handleDefault(c net.Conn) {
	body := `
		<!DOCTYPE html>
		<html lang="en">
		<head>
			<meta charset="UTF-8">
			<title>default</title>
		</head>
		<body>
			<h1>"default"</h1>
		</body>
		</html>
	`
	handleResponseHeaders(c, body)
}

func handleResponseHeaders(c net.Conn, body string) {
	io.WriteString(c, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(c, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(c, "Content-Type: text/html\r\n")
	io.WriteString(c, "\r\n")
	io.WriteString(c, body)
}
