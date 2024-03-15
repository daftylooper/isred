package Queue

import "fmt"

type Node struct {
	value string
	next  *Node
}

func MakeNode(value string) *Node {
	return &Node{value: value, next: nil}
}

func Enqueue(head *Node, value string) (*Node, error) {
	if head == nil {
		head = MakeNode(value)
	} else {
		var temp *Node = head
		for temp.next != nil {
			temp = temp.next
		}
		temp.next = MakeNode(value)
	}
	return head, nil
}

func Dequeue(head *Node) (*Node, string, error) {
	if head == nil {
		return head, "", fmt.Errorf("queue empty, cannot dequeue")
	} else {
		var dequeued *Node = head

		head = head.next
		dequeued.next = nil
		var dequeuedstr string = dequeued.value
		dequeued = nil

		return head, dequeuedstr, nil
	}
}

func DebugQueue(head *Node) error {
	if head == nil {
		return fmt.Errorf("qeueu empty!")
	} else {
		var temp *Node = head
		for temp.next != nil {
			fmt.Printf("->%s", temp.value)
			temp = temp.next
		}
		fmt.Printf("->%s\n", temp.value)
	}

	return nil
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
