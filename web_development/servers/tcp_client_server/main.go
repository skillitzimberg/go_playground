package main

import (
	"github.com/skillitzimberg/go_playground/web_development/servers/tcp_client_server/client"
	"github.com/skillitzimberg/go_playground/web_development/servers/tcp_client_server/server"
)

func main() {
	server.Response()

	for i := 0; i < 10; i++ {
		client.Request(i)
	}
}
