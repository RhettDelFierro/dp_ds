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