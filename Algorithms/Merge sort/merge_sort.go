package main

import "fmt"

func merge(arr1 []int, arr2 []int) []int {
	// Create a new array to store the merged elements.
	mergedArr := make([]int, len(arr1)+len(arr2))

	// Indices for the two input arrays.
	i := 0
	j := 0

	// Iterate through the two arrays and merge the elements.
	for k := 0; k < len(mergedArr); k++ {
		if i == len(arr1) {
			// If we have reached the end of arr1, just append the remaining elements from arr2.
			mergedArr[k] = arr2[j]
			j++
		} else if j == len(arr2) {
			// If we have reached the end of arr2, just append the remaining elements from arr1.
			mergedArr[k] = arr1[i]
			i++
		} else if arr1[i] < arr2[j] {
			// If the current element in arr1 is smaller, then append it to the merged array.
			mergedArr[k] = arr1[i]
			i++
		} else {
			// Otherwise, append the current element in arr2 to the merged array.
			mergedArr[k] = arr2[j]
			j++
		}
	}

	// Return the merged array.
	return mergedArr
}

func mergeSort(arr []int) []int {
	// If the array is empty or has only one element, then return it.
	if len(arr) <= 1 {
		return arr
	}

	// Split the array into two halves.
	mid := len(arr) / 2
	leftArr := mergeSort(arr[:mid])
	rightArr := mergeSort(arr[mid:])

	// Merge the two halves.
	return merge(leftArr, rightArr)
}

func main() {
	arr := []int{5, 2, 1, 6, 3, 4}
	fmt.Println("Original array:", arr)
	sortedArr := mergeSort(arr)
	fmt.Println("Sorted array:", sortedArr)
}
