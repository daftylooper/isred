package main

import (
	"fmt"
	"isred/Buffer"
	"isred/Engine"
	"isred/Engine/Cacher"
	"isred/Server"
	"log"
	"time"
)

func Getter(buf *Buffer.Buffer) {
	for {
		buf.GetCommand()
		// fmt.Println("->", str)
		buf.DebugBuffer()
		time.Sleep(1 * time.Second)
	}
}

func main() {
	replybuf, err := Buffer.InitialiseBuffer(-1)
	if err != nil {
		fmt.Println("Failure to Initialise Reply Buffer:", err)
	}

	server := Server.NewServer(":3000", replybuf)

	buf, err := Buffer.InitialiseBuffer(-1)
	if err != nil {
		fmt.Println("Failure to Initialise Buffer:", err)
	}

	kvs := Cacher.NewKeyValueStore()

	go func() {
		for msg := range server.GetMsgch() {
			err := buf.PushCommand(string(msg))
			if err != nil {
				fmt.Println("Couldn't Push:", err)
			}
			// buf.PersistBuffer("newfile")
		}
	}()

	// go Getter(buf)

	go Engine.EngineLoop(kvs, buf, replybuf)

	log.Fatal(server.Start())

}
