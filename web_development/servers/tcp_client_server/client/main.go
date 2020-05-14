package client

import (
	"fmt"
	"log"
	"net"
)

func Request(reqNum int) {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	fmt.Printf("The client is about to send request #%d\n.", reqNum)
	fmt.Fprintf(conn, "The client sends request: %d", reqNum)
}
