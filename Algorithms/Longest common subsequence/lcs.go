package main

import (
	"fmt"
)

// Max returns the maximum of two integers
func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// LCS returns the length and the longest common subsequence of two strings
func LCS(s1, s2 string) (int, string) {
	// Create a 2D array to store the lengths of common subsequences
	m := len(s1)
	n := len(s2)
	tab := make([][]int, m+1)
	for i := range tab {
		tab[i] = make([]int, n+1)
	}

	// Fill the array from bottom right to top left
	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if s1[i] == s2[j] {
				// If the characters match, add one to the diagonal value
				tab[i][j] = tab[i+1][j+1] + 1
			} else {
				// If not, take the maximum of the right and bottom values
				tab[i][j] = Max(tab[i+1][j], tab[i][j+1])
			}
		}
	}

	// The length of the longest common subsequence is at tab[0][0]
	length := tab[0][0]

	// To print the subsequence, start from tab[0][0] and move along the array
	subseq := make([]byte, length) // A slice to store the subsequence
	i := 0
	j := 0
	k := 0

	for i < m && j < n && k < length {
		if s1[i] == s2[j] {
			// If the characters match, append it to subseq and move diagonally
			subseq[k] = s1[i]
			i++
			j++
			k++
		} else if tab[i+1][j] >= tab[i][j+1] {
			// If not, move in the direction of the greater value
			i++
		} else {
			j++
		}
	}

	return length, string(subseq)
}

func main() {
	s1 := "AGGTAB"
	s2 := "GXTXAYB"
	length, subseq := LCS(s1, s2)
	fmt.Println("Length:", length)
	fmt.Println("Subsequence:", subseq)
}
