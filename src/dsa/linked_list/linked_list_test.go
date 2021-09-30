package linked_list

import (
	"fmt"
	"testing"
)

func TestSinglyLinkedList(t *testing.T) {
	l := SinglyLinkedList{}

	if !l.IsEmpty() {
		t.Errorf("Linked list should be empty.")
	}

	for i := 0; i < 10; i++ {
		l.Insert(i)
	}

	if l.size != 10 {
		t.Errorf("Linked list should contain 10 elements.")
	}

	if l.GetAt(5).el != 5 {
		t.Errorf("5 should be at the 5th position in the linked list.")
	}

	l.InsertAt(0, 1)
	if l.GetAt(2).el != 1 {
		t.Errorf("1 should be at the 2nd position in the linked list now.")
	}
	if l.size != 11 {
		t.Errorf("Linked list should contain 11 elements now")
	}

	if l.Search(40) != -1 {
		t.Errorf("-1 should be returned as the index for 40 in the linked list.")
	}

	if l.Search(5) != 6 {
		t.Errorf("6 should be returned as the index for 5 in the linked list.")
	}

	l.DeleteAt(1)
	l.DeleteAt(5)
	l.DeleteAt(2)

	if l.size != 8 {
		t.Errorf("Linked list currently contains 8 elements")
	}

	if l.GetAt(1).el != 1 {
		t.Errorf("1 should be returned.")
	}
	fmt.Println(l.ToArray())
}
