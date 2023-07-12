package main

import (
	"fmt"
)

// A Node represents a node in a Fibonacci heap.
type Node struct {
	// The key of the node.
	key int
	// The degree of the node.
	degree int
	// A flag indicating whether the node is marked.
	marked bool
	// The parent of the node.
	parent *Node
	// The children of the node.
	children []*Node
	// The next node in the root list.
	next *Node
	// The previous node in the root list.
	prev *Node
}

// A Heap represents a Fibonacci heap.
type Heap struct {
	// The minimum node in the heap.
	minimum *Node
	// The number of nodes in the heap.
	count int
}

// NewHeap creates a new Fibonacci heap.
func NewHeap() *Heap {
	return &Heap{
		minimum: nil,
		count:   0,
	}
}

// Insert inserts a new node into the heap.
func (h *Heap) Insert(node *Node) {
	// Set the node's degree to 0.
	node.degree = 0
	// Set the node's parent to nil.
	node.parent = nil
	// Set the node's marked flag to false.
	node.marked = false
	// Set the node's next and previous nodes to itself.
	node.next = node
	node.prev = node

	// Increment the heap's count.
	h.count++

	// If the heap is empty, then set the minimum node to the new node.
	if h.minimum == nil {
		h.minimum = node
	} else {
		// Otherwise, insert the new node into the root list.
		h.rootListInsert(node)
	}
}

// RemoveMin removes the minimum node from the heap and returns it.
func (h *Heap) RemoveMin() *Node {
	// If the heap is empty, then return nil.
	if h.minimum == nil {
		return nil
	}

	// Get the minimum node from the heap.
	min := h.minimum

	// Remove the minimum node from the root list.
	h.rootListRemove(min)

	// Decrement the heap's count.
	h.count--

	// Return the minimum node.
	return min
}

// rootListInsert inserts a node into the root list.
func (h *Heap) rootListInsert(node *Node) {
	// Insert the node into the root list after the minimum node.
	node.next = h.minimum
	node.prev = h.minimum.prev
	h.minimum.prev.next = node
	h.minimum.prev = node

	// If the node's key is less than the key of the minimum node, then set the minimum node to the new node.
	if node.key < h.minimum.key {
		h.minimum = node
	}
}

// rootListRemove removes a node from the root list.
func (h *Heap) rootListRemove(node *Node) {
	// If the node is the minimum node, then set the minimum node to the node's next node.
	if node == h.minimum {
		h.minimum = node.next
	}

	// Unlink the node from the root list.
	node.prev.next = node.next
	node.next.prev = node.prev

	// Set the node's next and previous nodes to nil.
	node.next = nil
	node.prev = nil
}

// consolidate merges any nodes that have the same degree.
func (h *Heap) consolidate() {
	// Create an array of all the roots in the heap.
	roots := make([]*Node, h.count)
	index := 0

	// Iterate over all the nodes in the heap and add them to the array.
	for node := h.minimum; node != nil; node = node.next {
		if node.marked {
			continue
		}

		roots[index] = node
		index++
	}

	// Iterate over the array and merge any nodes that have the same degree.
	for i := 0; i < index; i++ {
		for j := i + 1; j < index; j++ {
			if roots[i].degree == roots[j].degree {
				if roots[i].key > roots[j].key {
					// Swap the nodes so that roots[i] has the smaller key.
					tmp := roots[i]
					roots[i] = roots[j]
					roots[j] = tmp
				}

				// Merge the two nodes.
				h.merge(roots[i], roots[j])
			}
		}
	}

	// Reset the marked flags of all the nodes.
	for _, node := range roots {
		node.marked = false
	}
}

// merge merges two nodes in a Fibonacci heap.
func (h *Heap) merge(node1, node2 *Node) {
	// Set the parent of node2 to node1.
	node2.parent = node1

	// Increase the degree of node1.
	node1.degree++

	// If node2 has a smaller key than node1, then swap the two nodes.
	if node2.key < node1.key {
		tmp := node1
		node1 = node2
		node2 = tmp
	}

	// Append node2 to the list of children of node1.
	node1.children = append(node1.children, node2)
}

func main() {
	// Create a new Fibonacci heap.
	h := NewHeap()

	// Insert five nodes into the heap.
	h.Insert(&Node{key: 10})
	h.Insert(&Node{key: 5})
	h.Insert(&Node{key: 15})
	h.Insert(&Node{key: 20})
	h.Insert(&Node{key: 30})

	// Remove the minimum node from the heap and print its key.
	min := h.RemoveMin()
	fmt.Println(min.key) // 5

	// Remove the minimum node from the heap and print its key.
	min = h.RemoveMin()
	fmt.Println(min.key) // 10

	// Remove the minimum node from the heap and print its key.
	min = h.RemoveMin()
	fmt.Println(min.key) // 15
}
