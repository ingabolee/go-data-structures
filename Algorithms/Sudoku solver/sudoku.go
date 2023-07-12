package main

import "fmt"

func solveSudoku(board [][]int) bool {
	// If the board is already solved, return true.
	if isSolved(board) {
		return true
	}

	// Find the next empty cell in the board.
	row, col := findEmptyCell(board)
	if row == -1 {
		return true
	}

	// Try all possible values for the empty cell.
	for num := 1; num <= 9; num++ {
		// If the value is valid, place it in the empty cell and recursively solve the board.
		if isValid(board, row, col, num) {
			board[row][col] = num
			if solveSudoku(board) {
				return true
			}
		}
	}

	// If all possible values for the empty cell are invalid, backtrack.
	board[row][col] = 0
	return false
}

func isSolved(board [][]int) bool {
	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {
			if board[row][col] == 0 {
				return false
			}
		}
	}
	return true
}

func findEmptyCell(board [][]int) (int, int) {
	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {
			if board[row][col] == 0 {
				return row, col
			}
		}
	}
	return -1, -1
}

func isValid(board [][]int, row, col, num int) bool {
	// Check if the number is already in the row.
	for i := 0; i < 9; i++ {
		if board[row][i] == num {
			return false
		}
	}

	// Check if the number is already in the column.
	for i := 0; i < 9; i++ {
		if board[i][col] == num {
			return false
		}
	}

	// Check if the number is already in the 3x3 grid.
	rowStart := (row / 3) * 3
	colStart := (col / 3) * 3
	for i := rowStart; i < rowStart+3; i++ {
		for j := colStart; j < colStart+3; j++ {
			if board[i][j] == num {
				return false
			}
		}
	}

	return true
}

func main() {
	// Create a sudoku board.
	board := [][]int{
		{0, 3, 0, 0, 0, 0, 0, 0, 9},
		{9, 0, 0, 1, 8, 0, 0, 2, 0},
		{0, 0, 0, 0, 0, 6, 0, 0, 0},
		{0, 9, 0, 0, 0, 0, 0, 7, 0},
		{0, 0, 0, 2, 0, 0, 0, 0, 5},
		{0, 0, 0, 0, 7, 0, 0, 0, 4},
		{0, 0, 0, 0, 0, 3, 0, 1, 0},
		{2, 0, 0, 5, 0, 0, 0, 0, 6},
		{0, 8, 0, 0, 0, 0, 0, 0, 7},
	}

	// Solve the sudoku board.
	if solveSudoku(board) {
		// Print the solved sudoku board.
		for row := 0; row < 9; row++ {
			for col := 0; col < 9; col++ {
				fmt.Printf("%d ", board[row][col])
			}
			fmt.Println()
		}
	}
}
