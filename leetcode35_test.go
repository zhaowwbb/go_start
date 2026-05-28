package main

import (
	"fmt"
	"testing"
)

func TestSearchInsert(t *testing.T) {
	// Define our test case structure
	type testCase struct {
		nums     []int
		target   int
		expected int
	}

	// Slice of test cases
	tests := []testCase{
		{nums: []int{1, 3, 5, 6}, target: 5, expected: 2}, // Example 1
		{nums: []int{1, 3, 5, 6}, target: 2, expected: 1}, // Example 2
		{nums: []int{1, 3, 5, 6}, target: 7, expected: 4}, // Example 3
		{nums: []int{1, 3, 5, 6}, target: 0, expected: 0}, // Example 4
		{nums: []int{1}, target: 1, expected: 0},          // Single element match
		{nums: []int{1}, target: 0, expected: 0},          // Single element insert
	}

	// Print headers for clear output formatting
	fmt.Println("\n============================================================")
	fmt.Printf("%-25s | %-10s | %-10s | %s\n", "INPUT", "EXPECTED", "ACTUAL", "STATUS")
	fmt.Println("============================================================")

	for _, tc := range tests {
		// actual := SearchInsert(tc.nums, tc.target)
		actual := SearchInsertV2(tc.nums, tc.target)

		status := "✅ PASS"
		if actual != tc.expected {
			status = "❌ FAIL"
			// Flag the test framework that this test case failed
			t.Errorf("For nums=%v target=%d: expected %d, got %d", tc.nums, tc.target, tc.expected, actual)
		}

		// Format the input nicely for the console table
		inputStr := fmt.Sprintf("nums=%v, t=%d", tc.nums, tc.target)
		fmt.Printf("%-25s | %-10d | %-10d | %s\n", inputStr, tc.expected, actual, status)
	}
	// fmt.Println("============================================================\n")
}
