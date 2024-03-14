package main

import (
	"fmt"
	"isred/Buffer"
)

func main() {
	buf, err := Buffer.InitialiseBuffer(-1)
	if err != nil {
		fmt.Println("Initialise Buffer Error:", err)
	}

	buf, err = buf.PushCommand("hallo")
	buf, err = buf.PushCommand("is")
	buf, err = buf.PushCommand("this")
	buf, err = buf.PushCommand("a")
	buf, err = buf.PushCommand("buffer??!")

	// t := ""
	// buf, t = buf.GetCommand()
	// fmt.Println(t)

	buf.DebugBuffer()

	buf.PersistBuffer("buffer")

	buf, err = buf.ReadBuffer("buffer")
	if err != nil {
		fmt.Println("Error Reading Buffer:", err)
	}

	buf.DebugBuffer()
	buf, err = buf.PushCommand("NEWLY FORMED!")
	buf.DebugBuffer()
}
