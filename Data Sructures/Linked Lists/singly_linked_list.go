package main

import "fmt"

type Node struct {
	value int
	next  *Node
}

func main() {
	// Create a new node with the value 10.
	node1 := &Node{value: 10}

	// Create a new node with the value 20.
	node2 := &Node{value: 20}

	// Link the two nodes together.
	node1.next = node2

	// Print the value of the first node.
	fmt.Println(node1.value) // 10

	// Print the value of the second node.
	fmt.Println(node2.value) // 20
}
