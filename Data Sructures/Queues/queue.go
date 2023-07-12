package main

import (
	"fmt"
)

// Queue represents a queue data structure.
type Queue struct {
	// A pointer to the front of the queue.
	front *Node
	// A pointer to the rear of the queue.
	rear *Node
}

// Node represents a node in a queue.
type Node struct {
	// The value of the node.
	value int
	// A pointer to the next node in the queue.
	next *Node
}

// NewQueue creates a new queue.
func NewQueue() *Queue {
	// Initialize the front and rear fields to nil.
	q := &Queue{
		front: nil,
		rear:  nil,
	}

	// Return the queue.
	return q
}

// Enqueue inserts a new node into the queue.
func (q *Queue) Enqueue(value int) {
	// Create a new node with the given value.
	newNode := &Node{
		value: value,
		next:  nil,
	}

	// If the queue is empty, then set the front and rear to the new node.
	if q.front == nil {
		q.front = newNode
		q.rear = newNode
	} else {
		// Set the next pointer of the rear node to the new node.
		q.rear.next = newNode

		// Set the rear node to the new node.
		q.rear = newNode
	}
}

// Dequeue removes the front node from the queue.
func (q *Queue) Dequeue() int {
	// If the queue is empty, then return -1.
	if q.front == nil {
		return -1
	}

	// Save the value of the front node.
	value := q.front.value

	// Set the front node to the next node.
	q.front = q.front.next

	// If the queue is now empty, then set the rear to nil.
	if q.front == nil {
		q.rear = nil
	}

	// Return the value of the front node.
	return value
}

// main is the main function of the program.
func main() {
	// Create a new queue.
	q := NewQueue()

	// Enqueue three nodes into the queue.
	q.Enqueue(10)
	q.Enqueue(20)
	q.Enqueue(30)

	// Dequeue the three nodes from the queue.
	fmt.Println(q.Dequeue()) // 10
	fmt.Println(q.Dequeue()) // 20
	fmt.Println(q.Dequeue()) // 30
}
