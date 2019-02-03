package ds

import (
	"testing"
)

func TestQueueFromStack_Remove(t *testing.T) {
	queue := NewQueueFromStack()
	queue.Add([]string{"foo1"})
	queue.Add([]string{"foo2"})
	if queue.Peek() != "foo1" {
		t.Errorf("QueueFromStack.Add does not add correctly. Wanted: '%s', got: '%s'", "foo1", queue.Peek())
	}
	if queue.Peek() != "foo1" {
		t.Errorf("QueueFromStack.Add does not add correctly. Wanted: '%s', got: '%s'", "foo1", queue.Peek())
	}
	if queue.Remove() != "foo1" {
		t.Errorf("QueueFromStack.Remove does not remove correctly. Wanted: '%s', got: '%s'", "foo1", queue.Peek())
	}
	if queue.Remove() != "foo2" {
		t.Errorf("QueueFromStack.Remove does not remove correctly. Wanted: '%s', got: '%s'", "foo2", queue.Peek())
	}
}
