package main

import (
	"bytes"
	"testing"
)

func TestSolveSudoku(t *testing.T) {
	// 1. Define Input Board
	inputBoard := [][]byte{
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}

	// 2. Define Expected Output
	expectedOutput := [][]byte{
		{'5', '3', '4', '6', '7', '8', '9', '1', '2'},
		{'6', '7', '2', '1', '9', '5', '3', '4', '8'},
		{'1', '9', '8', '3', '4', '2', '5', '6', '7'},
		{'8', '5', '9', '7', '6', '1', '4', '2', '3'},
		{'4', '2', '6', '8', '5', '3', '7', '9', '1'},
		{'7', '1', '3', '9', '2', '4', '8', '5', '6'},
		{'9', '6', '1', '5', '3', '7', '2', '8', '4'},
		{'2', '8', '7', '4', '1', '9', '6', '3', '5'},
		{'3', '4', '5', '2', '8', '6', '1', '7', '9'},
	}

	// 3. Execute the function under test
	// solveSudoku(inputBoard)
	// solveSudokuV2(inputBoard)
	solveSudokuV3(inputBoard)

	// 4. Assert both matrices match element by element
	for i := 0; i < 9; i++ {
		// bytes.Equal is highly optimized for comparing slices of bytes
		if !bytes.Equal(inputBoard[i], expectedOutput[i]) {
			t.Errorf("Mismatch found at row %d.\nExpected: %s\nGot     : %s",
				i, string(expectedOutput[i]), string(inputBoard[i]))
		}
	}
}
