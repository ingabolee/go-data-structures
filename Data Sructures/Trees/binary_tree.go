package main

import "fmt"

// Define a struct to represent a binary tree node
type Node struct {
	value int   // value stored in the node
	left  *Node // pointer to the left child node
	right *Node // pointer to the right child node
}

// Insert method inserts a new node with the given value into the tree
func (n *Node) Insert(value int) {
	if value < n.value { // if value is less than current node's value
		if n.left == nil { // if left child node is nil, insert new node as left child
			n.left = &Node{value: value}
		} else { // else, recursively call Insert on left child node
			n.left.Insert(value)
		}
	} else { // if value is greater than or equal to current node's value
		if n.right == nil { // if right child node is nil, insert new node as right child
			n.right = &Node{value: value}
		} else { // else, recursively call Insert on right child node
			n.right.Insert(value)
		}
	}
}

// Contains method checks if the tree contains a node with the given value
func (n *Node) Contains(value int) bool {
	if n == nil { // if current node is nil, value not found
		return false
	}
	if n.value == value { // if current node's value matches value, value found
		return true
	} else if value < n.value { // if value is less than current node's value, check left child node
		return n.left.Contains(value)
	} else { // else, check right child node
		return n.right.Contains(value)
	}
}

// String method returns a string representation of the tree in-order
func (n *Node) String() string {
	if n == nil {
		return ""
	}
	str := n.left.String() // recursively get string for left subtree
	if str != "" {
		str += " "
	}
	str += fmt.Sprint(n.value) // add current node's value
	if n.right != nil {
		str += " "
		str += n.right.String() // recursively get string for right subtree
	}
	return str
}

func main() {
	// Create a new binary tree with the root node's value of 5
	tree := Node{value: 5}

	// Insert some nodes into the tree
	tree.Insert(3)
	tree.Insert(7)
	tree.Insert(1)
	tree.Insert(9)

	// Check if the tree contains nodes with certain values
	fmt.Println(tree.Contains(3)) // Output: true
	fmt.Println(tree.Contains(8)) // Output: false

	// Print the tree in-order
	fmt.Println(tree.String()) // Output: 1 3 5 7 9
}
