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

	buf = buf.PushCommand("hallo")
	buf = buf.PushCommand("is")
	buf = buf.PushCommand("this")
	buf = buf.PushCommand("a")
	buf = buf.PushCommand("buffer??!")

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
	buf = buf.PushCommand("NEWLY FORMED!")
	buf.DebugBuffer()
}
