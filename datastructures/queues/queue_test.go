package ds

import (
	"testing"
)

func NewQueue() *Queue {
	return &Queue{}
}

func TestAdd(t *testing.T) {
	queue := NewQueue()
	queue.Add([]string{"blah"})
	if len(*queue) != 1 {
		t.Errorf("Queue.Add does not add. Wanted: '%d', got: '%d'", 1, len(*queue))
	}
	queue.Add([]string{"blah"})
	if len(*queue) != 2 {
		t.Errorf("Queue.Add does not add. Wanted: '%d', got: '%d'", 2, len(*queue))
	}
}

func TestRemove(t *testing.T) {
	queue := NewQueue()
	queue.Add([]string{"blah"})
	queue.Add([]string{"foo"})
	queue.Remove()
	if len(*queue) != 1 {
		t.Errorf("Queue.Remove() does not remove. Wanted: '%d', got: '%d'", 1, len(*queue))
	}
}
