package main

import (
	"fmt"
	"isred/Server"
	"log"
)

func main() {
	server := Server.NewServer(":3000")

	go func() {
		for msg := range server.GetMsgch() {
			fmt.Printf("received message: %s\n", msg)
		}
	}()

	log.Fatal(server.Start())
}
