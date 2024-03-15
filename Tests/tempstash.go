// func main() {
// 	server := Server.NewServer(":3000")

// 	go func() {
// 		for msg := range server.GetMsgch() {
// 			fmt.Printf("received message: %s\n", msg)
// 		}
// 	}()

// 	log.Fatal(server.Start())
// }

// func main() {
// 	head := Queue.MakeNode("hallo")
// 	head = Queue.Enqueue(head, "this")
// 	head = Queue.Enqueue(head, "is")
// 	head = Queue.Enqueue(head, "me")
// 	head = Queue.Enqueue(head, "pranav")

// 	Queue.DebugQueue(head)

// 	t := ""
// 	head, t = Queue.Dequeue(head)
// 	fmt.Println(t)

// 	head = Queue.Enqueue(head, "ope, im back")

// 	Queue.DebugQueue(head)

// 	head, t = Queue.Dequeue(head)
// 	fmt.Println(t)

// 	Queue.DebugQueue(head)

// 	head, t = Queue.Dequeue(head)
// 	fmt.Println(t)
// 	head, t = Queue.Dequeue(head)
// 	fmt.Println(t)
// 	head, t = Queue.Dequeue(head)
// 	fmt.Println(t)
// 	head, t = Queue.Dequeue(head)
// 	fmt.Println(t)

// 	Queue.DebugQueue(head)

// 	head, t = Queue.Dequeue(head)
// 	fmt.Println(t)
// }

// func main() {
// 	buf, err := Buffer.InitialiseBuffer(-1)
// 	if err != nil {
// 		fmt.Println("Initialise Buffer Error:", err)
// 	}

// 	_ = buf.PushCommand("does")
// 	_ = buf.PushCommand("this")
// 	_ = buf.PushCommand("srsly")
// 	_ = buf.PushCommand("work")
// 	_ = buf.PushCommand("buffer??!")

// 	// t := ""
// 	// buf, t = buf.GetCommand()
// 	// fmt.Println(t)

// 	buf.DebugBuffer()

// 	buf.PersistBuffer("buffer")

// 	buf, err = buf.ReadBuffer("buffer")
// 	if err != nil {
// 		fmt.Println("Error Reading Buffer:", err)
// 	}

// 	buf.DebugBuffer()
// 	_ = buf.PushCommand("CYKAAA BLYATT")
// 	buf.DebugBuffer()
// }