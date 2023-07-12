package main

import (
	"fmt"
)

// Manacher returns the length and the longest palindromic substring of s
func Manacher(s string) (int, string) {
	// Insert "#" between each character and at the beginning and end
	t := "#"
	for _, c := range s {
		t += string(c) + "#"
	}

	// Create an array to store the length of the palindrome at each position
	p := make([]int, len(t))

	// Initialize the center and the right boundary of the current palindrome
	c := 0
	r := 0

	// Iterate over the modified string
	for i := 0; i < len(t); i++ {
		// Find the mirror position of i with respect to c
		mirror := 2*c - i

		// If i is within the current palindrome, use the mirror value as a lower bound
		if i < r {
			p[i] = min(r-i, p[mirror])
		}

		// Try to expand the palindrome centered at i
		for i-(p[i]+1) >= 0 && i+(p[i]+1) < len(t) && t[i-(p[i]+1)] == t[i+(p[i]+1)] {
			p[i]++
		}

		// Update the center and the right boundary if needed
		if i+p[i] > r {
			c = i
			r = i + p[i]
		}
	}

	// Find the maximum length and its index in p
	maxLen := 0
	maxIdx := 0
	for i, v := range p {
		if v > maxLen {
			maxLen = v
			maxIdx = i
		}
	}

	// Extract the longest palindromic substring from t
	lps := ""
	for i := maxIdx - maxLen; i <= maxIdx+maxLen; i++ {
		if t[i] != '#' {
			lps += string(t[i])
		}
	}

	return maxLen, lps
}

// min returns the minimum of two integers
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	s := "abcbabcbabcba"
	length, lps := Manacher(s)
	fmt.Println("Length:", length)
	fmt.Println("Longest palindromic substring:", lps)
}
