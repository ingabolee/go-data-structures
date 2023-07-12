package main

import (
	"fmt"
)

// Node represents a node in a doubly linked list.
type Node struct {
	// The value of the node.
	value int
	// A pointer to the previous node in the list.
	prev *Node
	// A pointer to the next node in the list.
	next *Node
}

// DoublyLinkedList represents a doubly linked list.
type DoublyLinkedList struct {
	// A pointer to the first node in the list.
	head *Node
	// A pointer to the last node in the list.
	tail *Node
}

// NewDoublyLinkedList creates a new doubly linked list.
func NewDoublyLinkedList() *DoublyLinkedList {
	// Initialize the head and tail fields to nil.
	l := &DoublyLinkedList{
		head: nil,
		tail: nil,
	}

	// Return the doubly linked list.
	return l
}

// Insert inserts a new node into the doubly linked list.
func (l *DoublyLinkedList) Insert(value int) {
	// Create a new node with the given value.
	newNode := &Node{
		value: value,
		prev:  nil,
		next:  nil,
	}

	// If the list is empty, then set the head and tail to the new node.
	if l.head == nil {
		l.head = newNode
		l.tail = newNode
	} else {
		// Set the next pointer of the tail node to the new node.
		l.tail.next = newNode

		// Set the prev pointer of the new node to the tail node.
		newNode.prev = l.tail

		// Set the tail node to the new node.
		l.tail = newNode
	}
}

// Print prints the doubly linked list.
func (l *DoublyLinkedList) Print() {
	// Start at the head node.
	node := l.head

	// Iterate through the list and print the value of each node.
	for node != nil {
		fmt.Println(node.value)
		node = node.next
	}
}

// main is the main function of the program.
func main() {
	// Create a new doubly linked list.
	list := NewDoublyLinkedList()

	// Insert three nodes into the list.
	list.Insert(10)
	list.Insert(20)
	list.Insert(30)

	// Print the list.
	list.Print()
}
