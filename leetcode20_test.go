package main

import (
	"fmt"
	"testing"
)

func TestIsValid(t *testing.T) {
	// Define test cases
	type testCase struct {
		input    string
		expected bool
	}

	tests := []testCase{
		{input: "()", expected: true},
		{input: "()[]{}", expected: true},
		{input: "(]", expected: false},
		{input: "([)]", expected: false},
		{input: "{[]}", expected: true},
		{input: "(", expected: false},
		{input: "]", expected: false},
		{input: "", expected: true},
	}

	successCount := 0
	failedCount := 0

	fmt.Println("\nExecuting Test Cases...")
	fmt.Println("----------------------------------------")

	for i, tc := range tests {
		// Call the implementation exactly once per test case
		// result := IsValid(tc.input)
		// result := IsValidV2(tc.input)
		result := IsValidV3(tc.input)

		if result == tc.expected {
			fmt.Printf("Test Case %d: PASSED\n", i+1)
			successCount++
		} else {
			fmt.Printf("Test Case %d: FAILED (Input: \"%s\" | Expected: %t | Got: %t)\n",
				i+1, tc.input, tc.expected, result)
			failedCount++
			t.Errorf("Test Case %d failed", i+1)
		}
	}

	fmt.Println("----------------------------------------")
	fmt.Printf("SUMMARY: %d passed, %d failed.\n\n", successCount, failedCount)
}
