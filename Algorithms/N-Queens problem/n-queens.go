package main

import "fmt"

func solveNQueens(n int) [][]string {
	// Initialize an empty board with n rows and n columns.
	board := make([][]byte, n)
	for i := range board {
		board[i] = make([]byte, n)
		for j := range board[i] {
			board[i][j] = '.'
		}
	}

	// Initialize the solutions array and start the backtracking algorithm.
	var solutions [][]string
	backtrack(board, 0, &solutions)
	return solutions
}

func backtrack(board [][]byte, row int, solutions *[][]string) {
	// If we have placed n queens, we have found a solution.
	if row == len(board) {
		*solutions = append(*solutions, createSolution(board))
		return
	}

	// Try to place a queen in each column of the current row.
	for col := 0; col < len(board); col++ {
		if isValid(board, row, col) {
			board[row][col] = 'Q'
			backtrack(board, row+1, solutions)
			board[row][col] = '.'
		}
	}
}

func isValid(board [][]byte, row, col int) bool {
	// Check if there is a queen in the same column.
	for i := 0; i < row; i++ {
		if board[i][col] == 'Q' {
			return false
		}
	}

	// Check if there is a queen in the same diagonal (upper left).
	for i, j := row-1, col-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if board[i][j] == 'Q' {
			return false
		}
	}

	// Check if there is a queen in the same diagonal (upper right).
	for i, j := row-1, col+1; i >= 0 && j < len(board); i, j = i-1, j+1 {
		if board[i][j] == 'Q' {
			return false
		}
	}

	// The placement is valid.
	return true
}

func createSolution(board [][]byte) []string {
	solution := make([]string, len(board))
	for i := range board {
		solution[i] = string(board[i])
	}
	return solution
}

func main() {
	n := 8
	solutions := solveNQueens(n)
	fmt.Printf("Solutions for n = %d:\n", n)
	for i, solution := range solutions {
		fmt.Printf("Solution %d:\n", i+1)
		for _, row := range solution {
			fmt.Println(row)
		}
		fmt.Println()
	}
}
