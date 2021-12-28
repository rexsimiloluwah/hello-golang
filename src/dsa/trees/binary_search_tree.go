package main

import (
	"fmt"
)

type Node struct {
	data  int
	left  *Node
	right *Node
}

type BST struct {
	root *Node
}

func main() {
	bst := BST{}
	bst.insert(10)
	bst.insert(7)
	bst.insert(9)
	bst.insert(14)
	bst.insert(11)
	bst.insert(20)
	bst.insert(22)
	bst.insert(8)
	bst.insert(1)
	bst.insert(6)

	fmt.Println("Inorder traversal result :-", InorderTraversal(bst.root))
}

func (bst *BST) insert(data int) {
	newNode := &Node{data, nil, nil}
	if bst.root == nil {
		bst.root = newNode
		return
	}

	currentNode := bst.root
	fmt.Println(currentNode.data)
	for currentNode != nil {
		if data < currentNode.data {
			if currentNode.left == nil {
				currentNode.left = &Node{data, nil, nil}
				return
			}
			currentNode = currentNode.left
		}
		if data > currentNode.data {
			if currentNode.right == nil {
				currentNode.right = &Node{data, nil, nil}
				return
			}
			currentNode = currentNode.right
		}
	}
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
