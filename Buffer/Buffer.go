package Buffer

import (
	"isred/Buffer/Queue"
	"os"
)

type Buffer struct {
	head          *Queue.Node
	replicationID int
	size          int
}

func InitialiseBuffer(replicationID int) *Buffer {
	if replicationID != -1 {
		//looks for saved state, if it is there, reload
		// &Buffer{head: initialised, replicationID: 1}
	} else {
		//else initialise an empty buffer
		var initialised *Queue.Node = Queue.MakeNode("nil")
		return &Buffer{head: initialised, replicationID: 0, size: 0}
	}
}

func (buf *Buffer) PersistBuffer() {
	//store buffer to disk
	f, err := os.Create("buffer")
	if err != nil {
		panic(err)
	}

	defer f.Close()

	var toWrite string = ""
	for command := range values {
		toWrite += string(command) + ";"
	}
	_, err = f.WriteString(toWrite)
	if err != nil {
		panic(err)
	}

	f.Sync()
}

func (buf *Buffer) PushCommand(command string) *Buffer {
	//enqueue with err
	buf.head = Queue.Enqueue(buf.head, command)
	return buf
}

func (buf *Buffer) GetCommand() (*Buffer, string) {
	//dequeue with val, err
	var command string
	buf.head, command = Queue.Dequeue(buf.head)
	return buf, command
}
