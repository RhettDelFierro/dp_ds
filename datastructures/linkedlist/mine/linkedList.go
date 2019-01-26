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

func (list *LinkedList) GetFirst() *Node { return list.Head }

func (list *LinkedList) GetLast() *Node {
	if list.Head == nil { return list.Head }

	currentNode := list.Head
	for currentNode != nil {
		if currentNode.Next == nil {
			return currentNode
		}
		currentNode = currentNode.Next
	}

	return nil
}

func (list *LinkedList) Clear() {
	list.Head = nil
}

func (list *LinkedList) RemoveFirst() {
	if list.Head == nil { return }
	list.Head = list.Head.Next
}
