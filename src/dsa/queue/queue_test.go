package queue

import (
	"testing"
)

func TestQueue(t *testing.T) {
	q := Queue{}

	if q.Size() != 0 {
		t.Errorf("Length of an empty queue should be 0.")
	}

	for i := 0; i < 10; i++ {
		q.Enqueue(i)
	}

	if q.Size() != 10 {
		t.Errorf("Length of the queue should be 10 after enqueueing 10 elements.")
	}

	if q.Peek() != 0 {
		t.Errorf("Current front element should be 0.")
	}

	for i := 0; i < 5; i++ {
		q.Dequeue()
	}

	if q.Size() != 5 {
		t.Errorf("Length of the queue should reduce to 5.")
	}

	for i := 0; i < 5; i++ {
		q.Dequeue()
	}

	if !q.IsEmpty() {
		t.Errorf("Queue should be empty.")
	}
}
