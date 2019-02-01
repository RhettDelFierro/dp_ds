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
	if list.Head == nil {
		return list.Head
	}

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
	if list.Head == nil {
		return
	}
	list.Head = list.Head.Next
}

func (list *LinkedList) RemoveLast() {
	if list.Head == nil {
		return
	}
	// if this is the only element in the list, remove it:
	if list.Head.Next == nil {
		list.Head = nil
	}

	previousNode := list.Head
	currentNode := list.Head.Next
	for currentNode.Next != nil {
		previousNode = currentNode
		currentNode = currentNode.Next
	}

	previousNode.Next = nil
}

func (list *LinkedList) InsertLast(data string) {
	last := list.GetLast()

	if last != nil {
		last.Next = &Node{Data: data}
	} else {
		list.Head = &Node{Data: data}
	}
}

func (list *LinkedList) GetAt(index int) *Node {
	counter := 0
	currentNode := list.Head

	for currentNode != nil {
		if index == counter {
			return currentNode
		}

		counter++
		currentNode = currentNode.Next
	}

	return nil
}

func (list *LinkedList) RemoveAt(index int) {
	if list.Head == nil {
		return
	}
	if index == 0 {
		list.Head = list.Head.Next
	}

	// check if out of range:
	if previousNode := list.GetAt(index - 1); previousNode == nil || previousNode.Next == nil {
		return
	} else {
		previousNode.Next = previousNode.Next.Next
	}
}

func (list *LinkedList) InsertAt(data string, index int) {
	if list.Head == nil || index == 0 {
		list.InsertFirst(data)
	}

	if previousNode := list.GetAt(index - 1); previousNode == nil {
		lastNode := list.GetLast()
		lastNode.Next = &Node{
			Data: data,
			Next: lastNode.Next,
		}
	} else {
		// connect the inserted node to the previous node's next:
		previousNode.Next = &Node{
			Data: data,
			Next: previousNode.Next,
		}
	}
}

func (list *LinkedList) ForEach(fn func(string) string) {
	currentNode := list.Head
	for currentNode != nil {
		currentNode.Data = fn(currentNode.Data)
		currentNode = currentNode.Next
	}
}
