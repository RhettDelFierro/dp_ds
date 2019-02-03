package ds

import (
	"testing"
)

func NewStack() *Stack {
	return &Stack{}
}

func TestAdd(t *testing.T) {
	stack := NewStack()
	stack.Add([]string{"blah"})
	if len(*stack) != 1 {
		t.Errorf("Stack.Add does not add. Wanted: '%d', got: '%d'", 1, len(*stack))
	}
	stack.Add([]string{"blah"})
	if len(*stack) != 2 {
		t.Errorf("Stack.Add does not add. Wanted: '%d', got: '%d'", 2, len(*stack))
	}
}

func TestRemove(t *testing.T) {
	stack := NewStack()
	stack.Add([]string{"blah"})
	stack.Add([]string{"foo"})
	stack.Remove()
	if len(*stack) != 1 {
		t.Errorf("Stack.Remove() does not remove. Wanted: '%d', got: '%d'", 1, len(*stack))
	}
}
