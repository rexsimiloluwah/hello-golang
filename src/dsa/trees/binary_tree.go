package main

import "fmt"

// Node in the Binary Tree Data Structure
type Node struct {
	data  interface{}
	left  *Node
	right *Node
}

// Binary Tree - A hierarchical linked data structure
type BinaryTree struct {
	root *Node
}

type Queue struct {
	data []*Node
	size int
}

func main() {
	bt := BinaryTree{}
	bt.root = &Node{10, nil, nil}
	bt.root.left = &Node{7, nil, nil}
	bt.root.right = &Node{11, nil, nil}
	bt.root.left.left = &Node{6, nil, nil}
	bt.root.left.left.left = &Node{1, nil, nil}
	bt.root.left.right = &Node{8, nil, nil}
	bt.root.left.right.right = &Node{9, nil, nil}
	bt.root.right.right = &Node{20, nil, nil}
	bt.root.right.right.left = &Node{14, nil, nil}
	bt.root.right.right.right = &Node{22, nil, nil}

	fmt.Println("Inorder traversal result: -", InorderTraversal(bt.root))
	fmt.Println("Preorder traversal result: -", PreorderTraversal(bt.root))
	fmt.Println("Postorder traversal result: -", PostorderTraversal(bt.root))
	fmt.Println("Levelorder traversal result: -", LevelorderTraversal((bt.root)))
}

// Traverse a binary tree using the in-order traversal algorithm (left-root-right) recursively
func InorderTraversal(rootNode *Node) []interface{} {
	// Base Case
	result := make([]interface{}, 0)
	if rootNode != nil {
		result = append(result, InorderTraversal(rootNode.left)...)
		result = append(result, rootNode.data)
		result = append(result, InorderTraversal(rootNode.right)...)
	}
	return result
}

// Traverse a binary tree using pre-order traversal algorithm (root-left-right) recursively
func PreorderTraversal(rootNode *Node) []interface{} {
	// Base Case
	result := make([]interface{}, 0)
	if rootNode != nil {
		result = append(result, rootNode.data)
		result = append(result, PreorderTraversal(rootNode.left)...)
		result = append(result, PreorderTraversal(rootNode.right)...)
	}
	return result
}

// Traverse a binary tree using a post-order traversal algorithm (left-right-root) recursively
func PostorderTraversal(rootNode *Node) []interface{} {
	// Base Case
	result := make([]interface{}, 0)
	if rootNode != nil {
		result = append(result, PostorderTraversal(rootNode.left)...)
		result = append(result, PostorderTraversal(rootNode.right)...)
		result = append(result, rootNode.data)
	}
	return result
}

// Level order traversal
func LevelorderTraversal(rootNode *Node) []interface{} {
	q := Queue{}
	result := make([]interface{}, 0)
	q.Enqueue(rootNode)
	result = append(result, rootNode.data)
	for q.size > 0 {
		for i := 0; i < q.size; i++ {
			currentNode := q.Dequeue()
			if currentNode.left != nil {
				q.Enqueue(currentNode.left)
				result = append(result, currentNode.left.data)
			}
			if currentNode.right != nil {
				q.Enqueue(currentNode.right)
				result = append(result, currentNode.right.data)
			}
		}
	}
	return result
}

func (q *Queue) Enqueue(data *Node) {
	q.data = append(q.data, data)
	q.size++
}

func (q *Queue) Dequeue() *Node {
	dequeued := q.data[0]
	q.data = q.data[1:]
	q.size--
	return dequeued
}
