package Buffer

import (
	"bufio"
	"fmt"
	"isred/Buffer/Queue"
	"os"
	"strings"
)

type Buffer struct {
	head          *Queue.Node
	replicationID int
	size          int
}

// func InitialiseBuffer(replicationID int) *Buffer {
// 	if replicationID != -1 {
// 		//looks for saved state, if it is there, reload

// 		// &Buffer{head: initialised, replicationID: 1}
// 	} else {
// 		//else initialise an empty buffer
// 		var initialised *Queue.Node = Queue.MakeNode("nil")
// 		return &Buffer{head: initialised, replicationID: 0, size: 0}
// 	}
// }

func (buf *Buffer) ReadBuffer() {
	// Open the file
	file, err := os.Open("buffer0")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Read the contents of the file
	scanner := bufio.NewScanner(file)
	var bufferValues string
	for scanner.Scan() {
		bufferValues = scanner.Text()
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	fmt.Println("-----> ", bufferValues)

	// Split the content by the delimiter ";"
	parts := strings.Split(bufferValues, ";")

	// Extract metadata and buffer values
	size := parts[0]
	replicationID := parts[1]
	buffer := parts[2:]

	// Print metadata and buffer values
	fmt.Println("Size:", size)
	fmt.Println("ReplicationID:", replicationID)
	fmt.Println("Buffer Values:", buffer)
}

func InitialiseBuffer(replicationID int) *Buffer {
	var initialised *Queue.Node = nil
	return &Buffer{head: initialised, replicationID: 0, size: 0}
}

func (buf *Buffer) PersistBuffer() {
	//store buffer to disk
	f, err := os.Create("buffer" + fmt.Sprintf("%d", buf.replicationID))
	if err != nil {
		panic(err)
	}

	defer f.Close()

	//generate string to write to file
	toWrite := fmt.Sprintf("%d", buf.replicationID) + ";" + fmt.Sprintf("%d", buf.size) + ";" + Queue.ConcatValues(buf.head)
	_, err = f.WriteString(toWrite)
	if err != nil {
		panic(err)
	}

	f.Sync()
}

func (buf *Buffer) PushCommand(command string) *Buffer {
	//enqueue with err
	buf.head = Queue.Enqueue(buf.head, command)
	buf.size += 1
	return buf
}

// iteratively gets next command from buffer queue
func (buf *Buffer) GetCommand() (*Buffer, string) {
	//dequeue with val, err
	command := ""
	buf.head, command = Queue.Dequeue(buf.head)
	buf.size -= 1
	return buf, command
}

func (buf *Buffer) DebugBuffer() {
	fmt.Printf("Replication ID: %d\n", buf.replicationID)
	fmt.Println("BUFFER:\n--------")
	Queue.DebugQueue(buf.head)
	fmt.Printf("Size: %d", buf.size)
}
