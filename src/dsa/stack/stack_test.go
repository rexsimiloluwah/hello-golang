package stack

import (
	"testing"
)

func TestStack(t *testing.T) {
	s := Stack{}

	if s.Size() != 0 {
		t.Errorf("Length of an empty stack should be 0.")
	}

	for i := 0; i < 10; i++ {
		s.Push(i)
	}

	if s.Size() != 10 {
		t.Errorf("Length of the stack should be 10 after pushing 10 elements.")
	}

	if s.Peek() != 9 {
		t.Errorf("Current top element should be 9.")
	}

	for i := 0; i < 5; i++ {
		s.Pop()
	}

	if s.Size() != 5 {
		t.Errorf("Length of the stack should reduce to 5.")
	}

	if s.Peek() != 4 {
		t.Errorf("Top element should be 4 currently.")
	}

	for i := 0; i < 5; i++ {
		s.Pop()
	}

	if !s.IsEmpty() {
		t.Errorf("Stack should be empty.")
	}
}
