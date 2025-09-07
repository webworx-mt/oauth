package main

import (
	"github.com/webworx-mt/oauth/server"
)

func main() {
	port := ":8080"
	server.StartServer(port)
}
