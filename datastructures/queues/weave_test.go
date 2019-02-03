package ds

import (
	"testing"
)

func TestWeave(t *testing.T) {
	queueOne := NewQueue()
	queueOne.Add([]string{"foo1"})
	queueOne.Add([]string{"foo2"})
	queueOne.Add([]string{"foo3"})
	queueOne.Add([]string{"foo4"})

	queueTwo := NewQueue()
	queueTwo.Add([]string{"bar1"})
	queueTwo.Add([]string{"bar2"})
	queueTwo.Add([]string{"bar3"})
	queueTwo.Add([]string{"bar4"})
	result := weave(queueOne, queueTwo)
	res1 := result.Remove()
	if res1 != "foo1" {
		t.Errorf("Queue.Add does not add. Wanted: '%s', got: '%s'", "foo1", res1)
	}
	res2 := result.Remove()
	if res2 != "bar1" {
		t.Errorf("Queue.Add does not add. Wanted: '%s', got: '%s'", "bar1", res2)
	}
}
