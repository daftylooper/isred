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