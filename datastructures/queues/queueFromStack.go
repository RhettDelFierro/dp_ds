package ds

import "github.com/RhettDelFierro/designPatterns/datastructures/stacks"

// see https://blog.golang.org/slices
type QueueFromStack struct {
	First  *ds.Stack
	Second *ds.Stack
}

func NewQueueFromStack() *QueueFromStack {
	return &QueueFromStack{
		First:  &ds.Stack{},
		Second: &ds.Stack{},
	}
}

func (q *QueueFromStack) Add(record []string) {
	*q.First = append(*q.First, record...)
}

func (q *QueueFromStack) Remove() string {
	for len(q.First.Peek()) > 0 {
		q.Second.Push([]string{q.First.Pop()})
	}
	record := q.Second.Pop()

	for len(q.Second.Peek()) > 0 {
		q.First.Push([]string{q.Second.Pop()})
	}
	return record
}

func (q *QueueFromStack) Peek() string {
	for len(q.First.Peek()) > 0 {
		q.Second.Push([]string{q.First.Pop()})
	}
	record := q.Second.Peek()

	for len(q.Second.Peek()) > 0 {
		q.First.Push([]string{q.Second.Pop()})
	}
	return record
}
