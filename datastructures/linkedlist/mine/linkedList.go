package ds

type Node struct {
	Data string
	Next *Node
}

type LinkedList struct {
	Head *Node
}

func (list *LinkedList) InsertFirst(data string) {
	node := &Node{
		Data: data,
		Next: list.Head,
	}
	list.Head = node
}

func (list *LinkedList) Size() int {
	counter := 0
	current := list.Head
	for current != nil {
		counter++
		current = current.Next
	}

	return counter
}