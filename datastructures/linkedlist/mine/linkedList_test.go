package ds

import (
	"testing"
)

func NewLinkedList() LinkedList {
	linkedList := LinkedList{}
	linkedList.InsertFirst("test")
	linkedList.InsertFirst("twotest")
	linkedList.InsertFirst("threetest")
	return linkedList
}

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

	if linkedList.Size() != 2 {
		t.Errorf("LinkedList.Size does returns wrong value. Wanted: '%d', got: '%d'", 2, linkedList.Size())
	}
}

func TestInsertSize(t *testing.T) {
	linkedList := LinkedList{}
	linkedList.InsertFirst("test")
	linkedList.InsertFirst("twotest")
	linkedList.InsertFirst("threetest")

	if linkedList.Size() != 3 {
		t.Errorf("LinkedList.Size does returns wrong value. Wanted: '%d', got: '%d'", 3, linkedList.Size())
	}
}

func TestGetFirst(t *testing.T) {
	linkedList := NewLinkedList()

	if linkedList.GetFirst().Data != "threetest" {
		t.Errorf("LinkedList.GetFirst returns wrong value. Wanted: '%s', got: '%s'", "threetest", linkedList.GetFirst().Data)
	}
}


func TestGetLast(t *testing.T) {
	linkedList := NewLinkedList()

	if linkedList.GetLast().Data != "test" {
		t.Errorf("LinkedList.GetFirst returns wrong value. Wanted: '%s', got: '%s'", "test", linkedList.GetLast().Data)
	}
}