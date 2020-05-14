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

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	request(conn)
}

func request(conn net.Conn) {
	i := 0
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
		if i == 0 { // All we care about here is the HTTP request start line.
			mux(conn, ln)
		}
		if ln == "" {
			break
		}
		i++
	}
}

func mux(conn net.Conn, reqLn string) {
	lineFields := strings.Fields(reqLn)

	m := lineFields[0]
	url := lineFields[1]
	req := fmt.Sprintf("%s %s", m, url)

	switch req {
	case "GET /":
		index(conn)
	case "GET /about":
		about(conn)
	case "GET /contact":
		contact(conn)
	case "GET /apply":
		apply(conn)
	case "POST /apply":
		applyProcess(conn)
	default:
		notFound(conn)
	}
}

func index(conn net.Conn) {
	html := formatBody("Home", false)

	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(html))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n") // CRLF (signals the end of response headers)
	fmt.Fprintf(conn, html)
}

func about(conn net.Conn) {
	html := formatBody("About", false)

	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(html))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n") // CRLF (signals the end of response headers)
	fmt.Fprintf(conn, html)
}

func contact(conn net.Conn) {
	html := formatBody("Contact", false)

	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(html))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n") // CRLF (signals the end of response headers)
	fmt.Fprintf(conn, html)
}

func apply(conn net.Conn) {
	html := formatBody("Apply", true)

	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(html))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n") // CRLF (signals the end of response headers)
	fmt.Fprintf(conn, html)
}

func applyProcess(conn net.Conn) {
	html := formatBody("Application Submitted", false)

	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(html))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n") // CRLF (signals the end of response headers)
	fmt.Fprintf(conn, html)
}

func notFound(conn net.Conn) {
	html := formatBody("Not Found", false)

	fmt.Fprint(conn, "HTTP/1.1 404 Not Found\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(html))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n") // CRLF (signals the end of response headers)
	fmt.Fprintf(conn, html)
}

func formatBody(title string, isForm bool) string {
	body := `<!DOCTYPE html><html lang="en"><head><meta charet="UTF-8"><title>%s</title></head><body>
	%s
	</body></html>`
	bodyWithForm := `<!DOCTYPE html><html lang="en"><head><meta charet="UTF-8"><title>%s</title></head><body>
	%s
	%s
	</body></html>`

	index := `<a href="/">index</a><br>`
	about := `<a href="/about">about</a><br>`
	contact := `<a href="/contact">contact</a><br>`
	apply := `<a href="/apply">apply</a><br>`
	form := `
	<form method="POST" action="/apply">
	<input type="submit" value="apply">
	</form>
	`

	nav := fmt.Sprintf("%s%s%s%s", index, about, contact, apply)
	html := fmt.Sprintf(body, title, nav)
	if isForm {
		html = fmt.Sprintf(bodyWithForm, title, nav, form)
	}
	return html
}
