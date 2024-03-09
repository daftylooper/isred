package Buffer

import (
	"bufio"
	"fmt"
	"isred/Buffer/Queue"
	"os"
	"strconv"
	"strings"
)

type Buffer struct {
	head          *Queue.Node
	replicationID int
	size          int
}

func InitialiseBuffer(replicationID int) (*Buffer, error) {
	if replicationID != -1 {
		var initialised *Queue.Node = nil
		return &Buffer{head: initialised, replicationID: 0, size: 0}, nil
	} else {
		// Open the file
		file, err := os.Open("buffer.buf")
		if err != nil {
			fmt.Println("Error opening file:", err)
			return nil, err
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
			return nil, err
		}

		// Split the content by the delimiter ";"
		parts := strings.Split(bufferValues, ";")

		replicationID, err := strconv.Atoi(parts[0])
		if err != nil {
			fmt.Println("Buffer File Parse Error: Couldn't Read Replication ID")
			return nil, err
		}
		size, err := strconv.Atoi(parts[1])
		if err != nil {
			fmt.Println("Buffer File Parse Error: Couldn't Read Size")
			return nil, err
		}
		var initialised *Queue.Node = nil
		buffer := parts[2 : len(parts)-1] //includes a blank val in parts array duie to strings.Split
		for _, bufvals := range buffer {
			initialised = Queue.Enqueue(initialised, bufvals)
		}

		return &Buffer{head: initialised, replicationID: replicationID, size: size}, nil
	}
}

func (buf *Buffer) PersistBuffer() {
	//store buffer to disk
	f, err := os.Create("buffer.buf")
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
