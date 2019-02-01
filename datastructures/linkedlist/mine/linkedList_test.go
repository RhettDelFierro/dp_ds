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
		t.Errorf("LinkedList.GetLast returns wrong value. Wanted: '%s', got: '%s'", "test", linkedList.GetLast().Data)
	}
}

func TestClear(t *testing.T) {
	linkedList := NewLinkedList()
	linkedList.Clear()
	if linkedList.Head != nil {
		t.Errorf("LinkedList.Clear does not clear list. Wanted: '%s', got: '%s'", (*Node)(nil), linkedList.Head)
	}
}

func TestRemoveFirst(t *testing.T) {
	linkedList := NewLinkedList()
	linkedList.RemoveFirst()
	if linkedList.Head.Data != "twotest" {
		t.Errorf("LinkedList.RemoveFirst() does not work. Wanted: '%s', got: '%s'", "twotest", linkedList.Head.Data)
	}
}

func TestRemoveLast(t *testing.T) {
	linkedList := NewLinkedList()
	linkedList.RemoveLast()
	if linkedList.GetLast().Data != "twotest" {
		t.Errorf("LinkedList.GetLast() does not work. Wanted: '%s', got: '%s'", "twotest", linkedList.GetLast().Data)
	}
}

func TestInsertLast(t *testing.T) {
	linkedList := NewLinkedList()
	linkedList.InsertLast("testminusone")
	if linkedList.GetLast().Data != "testminusone" {
		t.Errorf("LinkedList.InsertLast() does not work. Wanted: '%s', got: '%s'", "testminusone", linkedList.GetLast().Data)
	}
}

func TestGetAt(t *testing.T) {
	linkedList := NewLinkedList()
	linkedList.InsertLast("testminusone")
	two := linkedList.GetAt(1)

	if two.Data != "twotest" {
		t.Errorf("LinkedList.GetAt() does not work. Wanted: '%s', got: '%s'", "twotest", two.Data)
	}

	three := linkedList.GetAt(2)
	if three.Data != "test" {
		t.Errorf("LinkedList.GetAt() does not work. Wanted: '%s', got: '%s'", "test", three.Data)
	}
}

func TestRemoveAt(t *testing.T) {
	linkedList := NewLinkedList()
	linkedList.RemoveAt(1)
	test := linkedList.GetAt(1)

	if test.Data != "test" {
		t.Errorf("LinkedList.RemoveAt() does not work. Wanted: '%s', got: '%s'", "test", test.Data)
	}
}

func TestInsertAt(t *testing.T) {
	linkedList := NewLinkedList()
	linkedList.InsertAt("insertedtest", 2)
	four := linkedList.GetAt(3)

	if four.Data != "test" {
		t.Errorf("LinkedList.InsertAt() does not work. Wanted: '%s', got: '%s'", "test", four.Data)
	}

	inserted := linkedList.GetAt(2)
	if inserted.Data != "insertedtest" {
		t.Errorf("LinkedList.InsertAt() does not work. Wanted: '%s', got: '%s'", "insertedtest", inserted.Data)
	}
}

func TestForEach(t *testing.T) {
	linkedList := NewLinkedList()
	linkedList.ForEach(func(x string) string { return x + "!" })
	if linkedList.GetAt(1).Data != "twotest!" {
		t.Errorf("LinkedList.ForEach() does not work. Wanted: '%s', got: '%s'", "twotest!", linkedList.GetAt(1).Data)
	}
}
