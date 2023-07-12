package main

import (
	"fmt"
	"math/rand"
)

func QuickSort(arr []int) []int {
	// Base case: if the array has zero or one element, it is already sorted
	if len(arr) < 2 {
		return arr
	}
	// Choose a random pivot element
	pivotIndex := rand.Intn(len(arr))
	pivot := arr[pivotIndex]
	// Swap the pivot element with the last element of the array
	arr[pivotIndex], arr[len(arr)-1] = arr[len(arr)-1], arr[pivotIndex]
	// Initialize two pointers: i for the left subarray and j for the right subarray
	i, j := 0, len(arr)-2
	// Loop until i and j cross each other
	for i <= j {
		// If the element at i is smaller than or equal to the pivot, move i to the right
		if arr[i] <= pivot {
			i++
		} else {
			// Otherwise, swap the element at i with the element at j and move j to the left
			arr[i], arr[j] = arr[j], arr[i]
			j--
		}
	}
	// Swap the pivot element with the element at i, which is the first element larger than the pivot
	arr[i], arr[len(arr)-1] = arr[len(arr)-1], arr[i]
	// Recursively sort the left and right subarrays and concatenate them with the pivot element
	return append(append(QuickSort(arr[:i]), pivot), QuickSort(arr[i+1:])...)
}

// main function to test the quick sort algorithm
func main() {
	// Create a slice of random integers
	arr := []int{5, 3, 8, 2, 9, 1, 4, 7, 6}
	// Print the original slice
	fmt.Println("Original slice:", arr)
	// Sort the slice using quick sort
	sorted := QuickSort(arr)
	// Print the sorted slice
	fmt.Println("Sorted slice:", sorted)
}
