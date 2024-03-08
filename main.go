package main

import (
	"isred/Buffer"
)

func main() {
	buf := Buffer.InitialiseBuffer(-1)

	buf = buf.PushCommand("hallo")
	buf = buf.PushCommand("is")
	buf = buf.PushCommand("this")
	buf = buf.PushCommand("a")
	buf = buf.PushCommand("buffer??!")

	// t := ""
	// buf, t = buf.GetCommand()
	// fmt.Println(t)

	buf.DebugBuffer()

	buf.PersistBuffer()

	buf.ReadBuffer()
}
