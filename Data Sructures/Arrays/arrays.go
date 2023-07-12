package main

import (
	"fmt"
)

func main() {
	// Create a 1d array.
	array_1d := [5]int{1, 2, 3, 4, 5}

	// Print the array.
	fmt.Println(array_1d)

	// Access an element in the array.
	fmt.Println(array_1d[0])

	// Change an element in the array.
	array_1d[0] = 10
	fmt.Println(array_1d)

	// Create a 2d array.
	array_2d := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	// Print the 2d array.
	fmt.Println(array_2d)

	// Access an element in the 2d array.
	fmt.Println(array_2d[0][0])

	// Change an element in the 2d array.
	array_2d[0][0] = 10
	fmt.Println(array_2d)

	// Create a 3d array.
	array_3d := [][][]int{
		{
			{1, 2, 3},
			{4, 5, 6},
			{7, 8, 9},
		},
		{
			{10, 11, 12},
			{13, 14, 15},
			{16, 17, 18},
		},
	}

	// Print the 3d array.
	fmt.Println(array_3d)

	// Access an element in the 3d array.
	fmt.Println(array_3d[0][0][0])

	// Change an element in the 3d array.
	array_3d[0][0][0] = 10
	fmt.Println(array_3d)

	// Create a slice.
	slice_1d := array_1d[1:3]

	// Print the slice.
	fmt.Println(slice_1d)

	// Change an element in the slice.
	slice_1d[0] = 10
	fmt.Println(slice_1d)

	// Create a slice of a 2d array.
	slice_2d := array_2d[1:3]

	// Print the slice.
	fmt.Println(slice_2d)

	// Change an element in the slice.
	slice_2d[0][0] = 10
	fmt.Println(slice_2d)
}
