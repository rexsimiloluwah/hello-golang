package main

import (
	"fmt"
	"math"
)

type QueueNode struct {
	value    interface{}
	priority int
	next     *QueueNode
}

type PriorityQueue struct {
	head    *QueueNode
	maxSize int
	size    int
}

type Graph struct {
	numVertices     int
	adjacencyMatrix [][]float64
}

func main() {
	// p := PriorityQueue{}
	// p.SetMaxSize(10)
	// fmt.Println(p.IsEmpty())
	// p.Enqueue(1, 0)
	// s := p.Dequeue()
	// fmt.Println(s)
	// p.Enqueue(2, 0)
	// p.Display()
	// p.Enqueue(12, 5)
	// fmt.Println(p.Dequeue())
	// p.Display()
	// p.Enqueue(3, 1)
	// p.Enqueue(7, 3)
	// p.Enqueue(24, 0)
	// // fmt.Println(p.Dequeue().value)
	// // fmt.Println(p.Dequeue().value)
	// fmt.Println(p.Peek())
	// p.Display()
	// fmt.Println(p.IsEmpty())

	g := NewGraph(10)
	g.addEdge(0, 1, 4)
	g.addEdge(0, 6, 7)
	g.addEdge(1, 6, 11)
	g.addEdge(1, 7, 20)
	g.addEdge(1, 2, 9)
	g.addEdge(2, 3, 6)
	g.addEdge(2, 4, 2)
	g.addEdge(3, 4, 10)
	g.addEdge(3, 5, 5)
	g.addEdge(4, 5, 15)
	g.addEdge(4, 7, 1)
	g.addEdge(4, 8, 5)
	g.addEdge(5, 8, 12)
	g.addEdge(6, 7, 1)
	g.addEdge(7, 8, 3)
	fmt.Println(g.adjacencyMatrix)

	fmt.Println(g.FindShortestPaths(0))
}

func (p *PriorityQueue) IsEmpty() bool {
	return p.size == 0
}

func (p *PriorityQueue) SetMaxSize(maxSize int) {
	p.maxSize = maxSize
}

// Enqueue operation
func (p *PriorityQueue) Enqueue(el interface{}, priority int) {
	if p.size == p.maxSize {
		panic("Queue is full.")
	}
	newQueueNode := &QueueNode{}
	newQueueNode.value = el
	newQueueNode.priority = priority

	if p.IsEmpty() {
		p.head = newQueueNode
		p.size++
		return
	}

	ptr := p.head
	if priority < p.head.priority {
		p.head = newQueueNode
		p.head.next = ptr
		p.size++
	} else {
		for i := 0; i < p.size && ptr.next != nil; i++ {
			if ptr.next.priority <= priority {
				ptr = ptr.next
				fmt.Println(ptr.value)
			}
		}

		temp := ptr.next
		ptr.next, newQueueNode.next = newQueueNode, temp
		p.size++
	}
}

// Dequeue operation
func (p *PriorityQueue) Dequeue() QueueNode {
	if p.IsEmpty() {
		panic("Queue is empty.")
	}

	dequeued := *p.head
	if p.head.next != nil {
		p.head = p.head.next
	} else {
		p.head = nil
	}
	p.size--
	return dequeued
}

// Peek operation
func (p *PriorityQueue) Peek() QueueNode {
	if p.IsEmpty() {
		panic("Queue is empty.")
	}

	return *p.head
}

// Display operation
func (p *PriorityQueue) Display() {
	if p.IsEmpty() {
		panic("Queue is empty.")
	}

	arr := make([]interface{}, p.size)
	ptr := p.head
	for i := 0; i < p.size && ptr != nil; i++ {
		arr[i] = ptr.value
		ptr = ptr.next
	}

	fmt.Println("Priority Queue: ", arr)
}

// Instantiate a new graph
func NewGraph(numVertices int) Graph {
	newGraph := Graph{}
	newGraph.numVertices = numVertices
	newGraph.adjacencyMatrix = make([][]float64, numVertices)
	for i := 0; i < numVertices; i++ {
		arr := make([]float64, numVertices)
		for i := 0; i < numVertices; i++ {
			arr[i] = -1
		}

		newGraph.adjacencyMatrix[i] = arr
	}

	return newGraph
}

// Add a new edge to the graph by updating the adjacency matrix
func (p *Graph) addEdge(v1 int, v2 int, weight float64) {
	if v1 == v2 {
		panic("A Node cannot be connected to it self.")
	}
	p.adjacencyMatrix[v1][v2] = weight
	p.adjacencyMatrix[v2][v1] = weight
}

// Dijkstra algorithm implementation for a graph
func (p *Graph) FindShortestPaths(v int) map[int]float64 {
	// Create an array to store visited vertices
	var visited []int
	// Create a map that stores the vertice and the corresponding distance
	// The initial distance from the start vertex is initialized to infinity for all vertices except the start vertex which is initialized to 0
	m := make(map[int]float64)
	for i := 0; i < p.numVertices; i++ {
		m[i] = math.Inf(0)
	}
	m[v] = 0
	// Initialize the priority queue
	q := PriorityQueue{}
	q.SetMaxSize(p.numVertices)
	q.Enqueue(float64(v), 0) // The distance is the priority
	q.Display()
	for {
		fmt.Println(q.size)
		if q.size <= 0 {
			break
		}
		dequeued := q.Dequeue()
		fmt.Println(q.size)
		var distance float64
		distance = float64(dequeued.priority)
		currentVertex := dequeued.value.(float64)
		visited = append(visited, int(currentVertex))
		fmt.Println("Visited: ", visited)

		for i := 0; i < p.numVertices; i++ {
			// Iterate through all the neighbor vertices
			if p.adjacencyMatrix[int(currentVertex)][i] != -1 {
				distance = p.adjacencyMatrix[int(currentVertex)][i]
				if !contains(visited, i) {
					oldDistance := m[i]
					newDistance := m[int(currentVertex)] + distance
					// Check if newDistance is greater than oldDistance
					// If true, enqueue the queue with the newDistance, and update the distance for that vertice in the map of vertices and distances
					fmt.Printf("%f %T %f %T\n", oldDistance, oldDistance, newDistance, newDistance)
					if oldDistance > newDistance {
						q.Enqueue(float64(i), int(newDistance))
						m[i] = float64(newDistance)
					}
				}
			}
		}

		//q.Display()
	}

	fmt.Println(m)
	return m
}

// Utility function to check if a slice contains an element
func contains(s []int, el int) bool {
	var result bool = false
	for i := 0; i < len(s); i++ {
		if el == s[i] {
			result = true
			break
		}
	}

	return result
}
