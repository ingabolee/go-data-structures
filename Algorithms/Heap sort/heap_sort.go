package main

import (
	"fmt"
)

// Heap sort algorithm.
func heapSort(array []int) {
	// Build a max heap from the array.
	buildMaxHeap(array)

	// Iterate over the array, swapping the root element with the last element and then heapifying the remaining elements.
	for i := len(array) - 1; i >= 0; i-- {
		// Swap the root element with the last element.
		array[0], array[i] = array[i], array[0]

		// Heapify the remaining elements.
		heapify(array, 0, i-1)
	}
}

// Build a max heap from the array.
func buildMaxHeap(array []int) {
	for i := len(array)/2 - 1; i >= 0; i-- {
		heapify(array, i, len(array)-1)
	}
}

// Heapify the array starting at index `i` and ending at index `end`.
func heapify(array []int, i, end int) {
	// Get the left and right child of the current node.
	left := 2*i + 1
	right := 2*i + 2

	// Find the largest of the current node, its left child, and its right child.
	largest := i
	if left <= end && array[left] > array[largest] {
		largest = left
	}

	if right <= end && array[right] > array[largest] {
		largest = right
	}

	// If the largest element is not the current node, swap them and heapify the subtree rooted at the largest element.
	if largest != i {
		array[i], array[largest] = array[largest], array[i]
		heapify(array, largest, end)
	}
}

func main() {
	// Create an array of unsorted numbers.
	unsorted := []int{13, 48, 1, 4, -4, -3, 20, 0, 100, 99, 11, 3, 7}

	// Sort the array using heap sort.
	heapSort(unsorted)

	// Print the sorted array.
	fmt.Println("Sorted array:", unsorted)
}
