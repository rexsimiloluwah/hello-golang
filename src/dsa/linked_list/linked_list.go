package linked_list

import (
	"fmt"
)

// A Node basically stores a value and serves a pointer to another Node in the SinglyLinkedList
type Node struct {
	el   interface{}
	next *Node
}

// A SinglyLinkedList basically contains the head node
type SinglyLinkedList struct {
	head *Node
	size int
}

// TODO : Implement DoublyLinkedList

// Checks if SinglyLinkedList is empty
func (l *SinglyLinkedList) IsEmpty() bool {
	return l.size == 0
}

// Inserts an element into the end of the linked list
func (l *SinglyLinkedList) Insert(el interface{}) {
	// Create a new Node
	n := Node{}
	n.el = el
	if l.IsEmpty() {
		l.head = &n
		l.size++
		return
	}

	// If the linked list is not empty, iterate until we find the Node with a nil next pointer (essentially the end of the linked list)
	ptr := l.head
	for i := 0; i < l.size; i++ {
		if ptr.next == nil {
			ptr.next = &n
			l.size++
			return
		}

		ptr = ptr.next
	}
}

// Gets the element at a specific position in a linked list
func (l *SinglyLinkedList) GetAt(pos int) *Node {
	ptr := l.head
	if pos <= 0 {
		return l.head
	}

	if pos > l.size-1 {
		return nil
	} else {
		for i := 0; i < pos; i++ {
			ptr = ptr.next
		}

		return ptr
	}
}

// Insert an element at a specific position in a linked list
func (l *SinglyLinkedList) InsertAt(pos int, el interface{}) {
	n := Node{} // Create a new node
	n.el = el
	ptr := l.head
	if pos == 0 {
		l.head = &n
		l.head.next = ptr
		l.size++
	} else {
		// The node to be replaced
		currentNode := l.GetAt(pos)
		prevNode := l.GetAt(pos - 1)
		prevNode.next, n.next = &n, currentNode
		l.size++
	}
}

// Traverse through and Print a Linked list
func (l *SinglyLinkedList) Print() {
	if l.IsEmpty() {
		fmt.Println("Linked list is empty")
	}
	ptr := l.head
	for i := 0; i < l.size; i++ {
		fmt.Println(ptr.el)
		ptr = ptr.next
	}
}

// Search for a value in a linked list
// For the case when the linked list contains repeated elements, return the index of the one which appears first
func (l *SinglyLinkedList) Search(el interface{}) int {
	if l.IsEmpty() {
		panic("Linked list is empty")
	}
	ptr := l.head
	for i := 0; i < l.size; i++ {
		if ptr.el == el {
			return i
		}

		ptr = ptr.next
	}
	return -1
}

// Delete an element at a specific position from a linked list
func (l *SinglyLinkedList) DeleteAt(pos int) {
	if l.IsEmpty() {
		panic("Linked list is empty")
	}
	ptr := l.head
	if pos == 0 {
		l.head = ptr.next
		return
	}

	prevNode := l.GetAt(pos - 1)
	nextNode := l.GetAt(pos + 1)
	prevNode.next = nextNode
	l.size--
}

// Convert a linked list to an array
func (l *SinglyLinkedList) ToArray() []interface{} {
	ptr := l.head
	var arr []interface{}
	for i := 0; i < l.size; i++ {
		el := ptr.el
		arr = append(arr, el)
		ptr = ptr.next
	}
	return arr
}

// TODO: Reverse a Singly linked list
