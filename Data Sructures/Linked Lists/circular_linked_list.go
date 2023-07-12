package main

import (
	"fmt"
)

// Node represents a node in a circular linked list.
type Node struct {
	// The value of the node.
	value int
	// A pointer to the next node in the list.
	next *Node
}

// CircularLinkedList represents a circular linked list.
type CircularLinkedList struct {
	// A pointer to the first node in the list.
	head *Node
}

// NewCircularLinkedList creates a new circular linked list.
func NewCircularLinkedList() *CircularLinkedList {
	// Initialize the head field to nil.
	l := &CircularLinkedList{
		head: nil,
	}

	// Return the circular linked list.
	return l
}

// Insert inserts a new node into the circular linked list.
func (l *CircularLinkedList) Insert(value int) {
	// Create a new node with the given value.
	newNode := &Node{
		value: value,
		next:  nil,
	}

	// If the list is empty, then set the head and tail to the new node.
	if l.head == nil {
		l.head = newNode
		newNode.next = newNode
	} else {
		// Set the next pointer of the current tail node to the new node.
		l.head.next = newNode

		// Set the next pointer of the new node to the head node.
		newNode.next = l.head
	}
}

// Print prints the circular linked list.
func (l *CircularLinkedList) Print() {
	// Start at the head node.
	node := l.head

	// Iterate through the list and print the value of each node.
	for node != nil {
		fmt.Println(node.value)
		node = node.next

		// If the node is the head node, then break out of the loop.
		if node == l.head {
			break
		}
	}
}

// main is the main function of the program.
func main() {
	// Create a new circular linked list.
	list := NewCircularLinkedList()

	// Insert three nodes into the list.
	list.Insert(10)
	list.Insert(20)
	list.Insert(30)

	// Print the list.
	list.Print()
}
