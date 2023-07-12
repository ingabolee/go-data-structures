package main

import (
	"fmt"
)

// Define the structure of the binary heap
type BinaryHeap struct {
	heap []int
}

// Define the Insert method that inserts a new element into the binary heap
func (h *BinaryHeap) Insert(element int) {
	// Append the new element to the end of the heap
	h.heap = append(h.heap, element)

	// Get the index of the last element in the heap
	index := len(h.heap) - 1

	// While the parent of the new element is greater than the new element itself, swap them
	for index > 0 && h.heap[(index-1)/2] < h.heap[index] {
		h.heap[(index-1)/2], h.heap[index] = h.heap[index], h.heap[(index-1)/2]
		index = (index - 1) / 2
	}
}

// Define the ExtractMax method that extracts and returns the maximum element from the binary heap
func (h *BinaryHeap) ExtractMax() int {
	// Get the maximum element from the root of the heap
	max := h.heap[0]

	// Replace the root of the heap with the last element in the heap
	h.heap[0] = h.heap[len(h.heap)-1]
	h.heap = h.heap[:len(h.heap)-1]

	// Get the index of the root element
	index := 0

	// While either child of the root is greater than it, swap it with its greater child
	for (index*2)+1 < len(h.heap) {
		childIndex := (index * 2) + 1

		if (childIndex+1 < len(h.heap)) && (h.heap[childIndex+1] > h.heap[childIndex]) {
			childIndex++
		}

		if h.heap[index] < h.heap[childIndex] {
			h.heap[index], h.heap[childIndex] = h.heap[childIndex], h.heap[index]
			index = childIndex
		} else {
			break
		}
	}

	return max
}

func main() {

	// Create a new instance of the BinaryHeap struct with an empty slice as its heap field
	binaryHeap := &BinaryHeap{heap: []int{}}

	// Insert some elements into the binary heap
	binaryHeap.Insert(10)
	binaryHeap.Insert(20)
	binaryHeap.Insert(56)
	binaryHeap.Insert(30)

	// Extract and print out some elements from the binary heap
	fmt.Println(binaryHeap.ExtractMax())

}
