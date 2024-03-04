package Queue

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
	var dequeued *Node = head

	head = head.next
	dequeued.next = nil
	var dequeuedstr string = dequeued.value
	dequeued = nil

	return head, dequeuedstr
}
