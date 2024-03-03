package main

import (
	"isred/Server"
	"log"
)

func main() {
	server := Server.NewServer(":3000")
	log.Fatal(server.Start())
}
