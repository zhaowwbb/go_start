package main

import "fmt"

// IsValidSudoku checks if a 9x9 Sudoku board is valid.
func IsValidSudoku(board [][]string) bool {
	// A map of empty structs acts as a space-efficient Set in Go
	seen := make(map[string]struct{})

	for r := 0; r < 9; r++ {
		for c := 0; c < 9; c++ {
			number := board[r][c]

			// Skip empty cells
			if number != "." {
				// Construct unique string markers for row, col, and sub-box
				rowKey := fmt.Sprintf("%s in row %d", number, r)
				colKey := fmt.Sprintf("%s in col %d", number, c)
				boxKey := fmt.Sprintf("%s in box %d-%d", number, r/3, c/3)

				// Check if any of these markers have been seen before
				if _, exists := seen[rowKey]; exists {
					return false
				}
				if _, exists := seen[colKey]; exists {
					return false
				}
				if _, exists := seen[boxKey]; exists {
					return false
				}

				// Record the markers in our map set
				seen[rowKey] = struct{}{}
				seen[colKey] = struct{}{}
				seen[boxKey] = struct{}{}
			}
		}
	}

	return true
}

func IsValidSudokuV2(board [][]string) bool {
	seen := make(map[string]struct{})

	for r := 0; r < 9; r++ {
		for c := 0; c < 9; c++ {
			number := board[r][c]
			if number != "." {
				rowKey := fmt.Sprintf("%s row key %d", number, r)
				colKey := fmt.Sprintf("%s col key %d", number, c)
				boxKey := fmt.Sprintf("%s box key %d-%d", number, r/3, c/3)

				if _, exists := seen[rowKey]; exists {
					return false
				}
				if _, exists := seen[colKey]; exists {
					return false
				}
				if _, exists := seen[boxKey]; exists {
					return false
				}

				seen[rowKey] = struct{}{}
				seen[colKey] = struct{}{}
				seen[boxKey] = struct{}{}
			}
		}
	}
	return true
}

func IsValidSudokuV3(board [][]string) bool {
	seen := make(map[string]struct{})

	for r := 0; r < 9; r++ {
		for c := 0; c < 9; c++ {
			number := board[r][c]
			if number != "." {
				rowKey := fmt.Sprintf("%s in row %d", number, r)
				colKey := fmt.Sprintf("%s in col %d", number, c)
				boxKey := fmt.Sprintf("%s in box %d-%d", number, r/3, c/3)

				if _, exists := seen[rowKey]; exists {
					return false
				}
				if _, exists := seen[colKey]; exists {
					return false
				}
				if _, exists := seen[boxKey]; exists {
					return false
				}

				seen[rowKey] = struct{}{}
				seen[colKey] = struct{}{}
				seen[boxKey] = struct{}{}
			}
		}
	}
	return true
}
