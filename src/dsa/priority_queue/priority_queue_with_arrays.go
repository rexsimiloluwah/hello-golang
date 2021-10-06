// package main

// import "fmt"

// type QueueElement struct {
// 	value    interface{}
// 	priority int
// }
// type PriorityQueue struct {
// 	maxSize int
// 	data    []QueueElement
// 	size    int
// }

// func main() {
// 	p := PriorityQueue{}
// 	p.SetMaxSize(10)
// 	fmt.Println(p.IsEmpty())
// 	p.Enqueue(1, 0)
// 	p.Enqueue(2, 0)
// 	p.Display()
// 	p.Enqueue(12, 5)
// 	p.Display()
// 	p.Enqueue(3, 1)
// 	p.Enqueue(7, 3)
// 	p.Enqueue(4, 1)
// 	p.Enqueue(24, 0)
// 	// fmt.Println(p.Dequeue().value)
// 	// fmt.Println(p.Dequeue().value)
// 	fmt.Println(p.Peek())
// 	p.Display()
// }

// func (p *PriorityQueue) SetMaxSize(maxSize int) {
// 	p.maxSize = maxSize
// }

// func (p *PriorityQueue) IsEmpty() bool {
// 	return p.size == 0
// }

// // Adds a new element to the priority queue in its exact location
// func (p *PriorityQueue) Enqueue(el interface{}, priority int) {
// 	newElement := QueueElement{el, priority}
// 	if p.size == p.maxSize {
// 		panic("Queue has reached its max size limit.")
// 	}

// 	if p.IsEmpty() {
// 		p.data = append(p.data, newElement)
// 		p.size++
// 		return
// 	}

// 	p.data = append(p.data, newElement)
// 	i := p.size - 1
// 	for i >= 0 {
// 		if p.data[i].priority > priority {
// 			p.data[i+1] = p.data[i]
// 			i--
// 		} else {
// 			break
// 		}
// 	}
// 	p.data[i+1] = newElement
// 	p.size++
// }

// // Returns and Removes the first element of the priority queue
// func (p *PriorityQueue) Dequeue() QueueElement {
// 	if p.IsEmpty() {
// 		panic("Queue is empty.")
// 	}
// 	dequeued := p.data[0]

// 	p.data = p.data[1:]
// 	p.size--
// 	return dequeued
// }

// // Returns the first element in the queue without modifying the queue
// func (p *PriorityQueue) Peek() interface{} {
// 	if p.IsEmpty() {
// 		panic("Queue is empty.")
// 	}

// 	return p.data[0]
// }

// // Display the elements of the queue in an array form
// func (p *PriorityQueue) Display() {
// 	if p.IsEmpty() {
// 		panic("Queue is empty.")
// 	}

// 	arr := make([]interface{}, p.size)
// 	i := 0
// 	for i < p.size {
// 		arr[i] = p.data[i].value
// 		i++
// 	}

// 	fmt.Println("Priority Queue elements: ", arr)
// }
