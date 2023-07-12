package main

import (
	"fmt"
	"math"
)

// Define a struct to represent a node in the AVL tree
type Node struct {
	value  int   // value stored in the node
	height int   // height of the node (max height of its subtrees + 1)
	left   *Node // pointer to the left child node
	right  *Node // pointer to the right child node
}

// getHeight function returns the height of a node
func getHeight(n *Node) int {
	if n == nil {
		return 0
	}
	return n.height
}

// getBalance function returns the balance factor of a node
func getBalance(n *Node) int {
	if n == nil {
		return 0
	}
	return getHeight(n.left) - getHeight(n.right)
}

// rotateLeft function performs a left rotation on a node
func rotateLeft(n *Node) *Node {
	r := n.right
	rl := r.left

	r.left = n
	n.right = rl

	n.height = 1 + int(math.Max(float64(getHeight(n.left)), float64(getHeight(n.right))))
	r.height = 1 + int(math.Max(float64(getHeight(r.left)), float64(getHeight(r.right))))

	return r
}

// rotateRight function performs a right rotation on a node
func rotateRight(n *Node) *Node {
	l := n.left
	lr := l.right

	l.right = n
	n.left = lr

	n.height = 1 + int(math.Max(float64(getHeight(n.left)), float64(getHeight(n.right))))
	l.height = 1 + int(math.Max(float64(getHeight(l.left)), float64(getHeight(l.right))))

	return l
}

// insert function inserts a new node with the given value into the tree
func insert(n *Node, value int) *Node {
	if n == nil {
		return &Node{value: value, height: 1}
	}

	if value < n.value {
		n.left = insert(n.left, value)
	} else if value > n.value {
		n.right = insert(n.right, value)
	} else {
		return n
	}

	n.height = 1 + int(math.Max(float64(getHeight(n.left)), float64(getHeight(n.right))))

	balance := getBalance(n)

	if balance > 1 && value < n.left.value {
		return rotateRight(n)
	}

	if balance < -1 && value > n.right.value {
		return rotateLeft(n)
	}

	if balance > 1 && value > n.left.value {
		n.left = rotateLeft(n.left)
		return rotateRight(n)
	}

	if balance < -1 && value < n.right.value {
		n.right = rotateRight(n.right)
		return rotateLeft(n)
	}

	return n
}

// traverseInOrder function traverses the tree in-order and returns a slice of values
func traverseInOrder(n *Node) []int {
	if n == nil {
		return []int{}
	}
	left := traverseInOrder(n.left)
	right := traverseInOrder(n.right)
	return append(append(left, n.value), right...)
}

func main() {
	// Create a new AVL tree
	var tree *Node

	// Insert some values into the tree
	tree = insert(tree, 10)
	tree = insert(tree, 20)
	tree = insert(tree, 30)
	tree = insert(tree, 40)
	tree = insert(tree, 50)
	tree = insert(tree, 25)

	// Traverse the tree in-order and print the values
	fmt.Println(traverseInOrder(tree)) // Output: [10 20 25 30 40 50]
}
