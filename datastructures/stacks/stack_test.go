package ds

import (
	"testing"
)

func NewStack() *Stack {
	return &Stack{}
}

func TestPush(t *testing.T) {
	stack := NewStack()
	stack.Push([]string{"blah"})
	if len(*stack) != 1 {
		t.Errorf("Stack.Push does not add. Wanted: '%d', got: '%d'", 1, len(*stack))
	}
	stack.Push([]string{"blah"})
	if len(*stack) != 2 {
		t.Errorf("Stack.Push does not add. Wanted: '%d', got: '%d'", 2, len(*stack))
	}
}

func TestPop(t *testing.T) {
	stack := NewStack()
	stack.Push([]string{"blah"})
	stack.Push([]string{"foo"})
	stack.Pop()
	if len(*stack) != 1 {
		t.Errorf("Stack.Pop() does not remove. Wanted: '%d', got: '%d'", 1, len(*stack))
	}
}
