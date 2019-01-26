package ds

import (
	"testing"
)

func TestInsertFirst(t *testing.T) {
	linkedList := LinkedList{}
	linkedList.InsertFirst("test")
	if linkedList.Head.Data != "test" {
		t.Errorf("LinkedList.InsertFirst does not store head. Wanted: '%s', got: '%s'", "test", linkedList.Head.Data)
	}

	if linkedList.Head.Next != (nil) {
		t.Errorf("Empty Linked List with LinkedList.InsertFirst has Next data.")
	}

	linkedList.InsertFirst("twotest")

	if linkedList.Head.Data != "twotest" {
		t.Errorf("LinkedList.InsertFirst does not store head. Wanted: '%s', got: '%s'", "twotest", linkedList.Head.Data)
	}

	if linkedList.Head.Next == (nil) {
		t.Errorf("Linked List did not create a new node.")
	}

	if linkedList.Head.Next.Data != "test" {
		t.Errorf("LinkedList.InsertFirst does not push head data. Wanted: '%s', got: '%s'", "test", linkedList.Head.Next.Data)
	}
}

