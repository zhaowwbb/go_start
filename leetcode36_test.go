package main

import (
	"fmt"
	"testing"
)

func TestIsValidSudoku(t *testing.T) {
	validBoard := [][]string{
		{"5", "3", ".", ".", "7", ".", ".", ".", "."},
		{"6", ".", ".", "1", "9", "5", ".", ".", "."},
		{".", "9", "8", ".", ".", ".", ".", "6", "."},
		{"8", ".", ".", ".", "6", ".", ".", ".", "3"},
		{"4", ".", ".", "8", ".", "3", ".", ".", "1"},
		{"7", ".", ".", ".", "2", ".", ".", ".", "6"},
		{".", "6", ".", ".", ".", ".", "2", "8", "."},
		{".", ".", ".", "4", "1", "9", ".", ".", "5"},
		{".", ".", ".", ".", "8", ".", ".", "7", "9"},
	}

	invalidBoard := [][]string{
		{"8", "3", ".", ".", "7", ".", ".", ".", "."}, // Changed top-left to '8'
		{"6", ".", ".", "1", "9", "5", ".", ".", "."},
		{".", "9", "8", ".", ".", ".", ".", "6", "."},
		{"8", ".", ".", ".", "6", ".", ".", ".", "3"},
		{"4", ".", ".", "8", ".", "3", ".", ".", "1"},
		{"7", ".", ".", ".", "2", ".", ".", ".", "6"},
		{".", "6", ".", ".", ".", ".", "2", "8", "."},
		{".", ".", ".", "4", "1", "9", ".", ".", "5"},
		{".", ".", ".", ".", "8", ".", ".", "7", "9"},
	}

	tests := []struct {
		name     string
		board    [][]string
		expected bool
	}{
		{name: "Example 1 (Valid Board)  ", board: validBoard, expected: true},
		{name: "Example 2 (Invalid Board)", board: invalidBoard, expected: false},
	}

	fmt.Println("\n============================================================")
	fmt.Printf("%-25s | %-10s | %-10s | %s\n", "TEST CASE", "EXPECTED", "ACTUAL", "STATUS")
	fmt.Println("============================================================")

	for _, tc := range tests {
		// actual := IsValidSudoku(tc.board)
		// actual := IsValidSudokuV2(tc.board)
		actual := IsValidSudokuV3(tc.board)
		status := "✅ PASS"

		if actual != tc.expected {
			status = "❌ FAIL"
			t.Errorf("%s failed: expected %v, got %v", tc.name, tc.expected, actual)
		}

		fmt.Printf("%-25s | %-10t | %-10t | %s\n", tc.name, tc.expected, actual, status)
	}
	// fmt.Println("============================================================\n")
}
