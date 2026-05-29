package main

// solveSudoku is the entry point that modifies the board in-place.
func solveSudoku(board [][]byte) {
	if board == nil || len(board) == 0 {
		return
	}
	solve(board)
}

// solve handles the recursive backtracking algorithm.
func solve(board [][]byte) bool {
	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {

			// Find an empty cell
			if board[row][col] == '.' {

				// Try digits '1' through '9'
				for c := byte('1'); c <= byte('9'); c++ {
					if isValid(board, row, col, c) {
						board[row][col] = c // Tentative assignment

						// Recursively attempt to solve the rest of the board
						if solve(board) {
							return true
						}

						board[row][col] = '.' // Backtrack: undo our choice
					}
				}
				return false // Triggers backtracking to the previous level
			}
		}
	}
	return true // Board is fully solved
}

// isValid checks if a character placement satisfies Sudoku constraints.
func isValid(board [][]byte, row, col int, c byte) bool {
	for i := 0; i < 9; i++ {
		// 1. Check Row constraint
		if board[row][i] == c {
			return false
		}

		// 2. Check Column constraint
		if board[i][col] == c {
			return false
		}

		// 3. Check 3x3 Box constraint
		boxRow := 3*(row/3) + i/3
		boxCol := 3*(col/3) + i%3
		if board[boxRow][boxCol] == c {
			return false
		}
	}
	return true
}

func solveSudokuV2(board [][]byte) {
	if len(board) == 0 {
		return
	}
	solveV2(board)
}

func solveV2(board [][]byte) bool {
	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {
			if board[row][col] == '.' {
				for num := byte('1'); num <= byte('9'); num++ {
					if isValidV2(board, row, col, num) {
						board[row][col] = num
						if solveV2(board) {
							return true
						}
						board[row][col] = '.'
					}
				}
				return false
			}
		}
	}
	return true
}

func isValidV2(board [][]byte, row int, col int, number byte) bool {
	for i := 0; i < 9; i++ {
		if board[row][i] == number {
			return false
		}
		if board[i][col] == number {
			return false
		}
	}

	boxRow := (row / 3) * 3
	boxCol := (col / 3) * 3
	for r := boxRow; r < boxRow+3; r++ {
		for c := boxCol; c < boxCol+3; c++ {
			if board[r][c] == number {
				return false
			}
		}
	}

	return true
}

func solveSudokuV3(board [][]byte) {
	if len(board) == 0 {
		return
	}
	solveV3(board)
}

func solveV3(board [][]byte) bool {
	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {
			if board[row][col] == '.' {
				for num := byte('1'); num <= byte('9'); num++ {
					if isValidV3(board, row, col, num) {
						board[row][col] = num
						if solveV3(board) {
							return true
						}
						board[row][col] = '.'
					}
				}
				return false
			}
		}
	}
	return true
}

func isValidV3(board [][]byte, row int, col int, num byte) bool {
	for i := 0; i < 9; i++ {
		if board[row][i] == num {
			return false
		}
		if board[i][col] == num {
			return false
		}
	}
	boxRow := (row / 3) * 3
	boxCol := (col / 3) * 3
	for r := boxRow; r < boxRow+3; r++ {
		for c := boxCol; c < boxCol+3; c++ {
			if board[r][c] == num {
				return false
			}
		}
	}
	return true
}
