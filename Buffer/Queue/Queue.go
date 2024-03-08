package Queue

import "fmt"

type Node struct {
	value string
	next  *Node
}

func MakeNode(value string) *Node {
	return &Node{value: value, next: nil}
}

func Enqueue(head *Node, value string) *Node {
	if head == nil {
		head = MakeNode(value)
	} else {
		var temp *Node = head
		for temp.next != nil {
			temp = temp.next
		}
		temp.next = MakeNode(value)
	}
	return head
}

func Dequeue(head *Node) (*Node, string) {
	if head == nil {
		fmt.Println("queue empty, cannot dequeue")
		return head, "-1"
	} else {
		var dequeued *Node = head

		head = head.next
		dequeued.next = nil
		var dequeuedstr string = dequeued.value
		dequeued = nil

		return head, dequeuedstr
	}
}

func DebugQueue(head *Node) {
	if head == nil {
		fmt.Println("queue empty!")
	} else {
		var temp *Node = head
		for temp.next != nil {
			fmt.Printf("->%s", temp.value)
			temp = temp.next
		}
		fmt.Printf("->%s\n", temp.value)
	}
}

func ConcatValues(head *Node) string {
	temp := head
	concat := ""
	for temp.next != nil {
		concat += temp.value + ";"
		temp = temp.next
	}
	concat += temp.value + ";"

	return concat
}
