package main

import (
	"fmt"
	"testing"
)

// TestCase defines the structure for our test inputs and expectations
type TestCase struct {
	input       []int
	expectedK   int
	expectedArr []int
}

// equal checks if two integer slices are identical
func equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func TestRemoveDuplicates(t *testing.T) {
	testCases := []TestCase{
		{input: []int{1, 1, 2}, expectedK: 2, expectedArr: []int{1, 2}},
		{input: []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}, expectedK: 5, expectedArr: []int{0, 1, 2, 3, 4}},
		{input: []int{}, expectedK: 0, expectedArr: []int{}},
		{input: []int{1}, expectedK: 1, expectedArr: []int{1}},
		{input: []int{1, 1, 1, 1}, expectedK: 1, expectedArr: []int{1}},
		{input: []int{1, 2, 3, 4}, expectedK: 4, expectedArr: []int{1, 2, 3, 4}},
	}

	fmt.Println("Running test cases...")
	fmt.Println("--------------------------------------------------")

	passedCount := 0
	failedCount := 0

	for i, tc := range testCases {
		// Create a fresh copy of the input slice since RemoveDuplicates modifies it in-place
		numsInput := make([]int, len(tc.input))
		copy(numsInput, tc.input)

		// Call the implementation exactly once per test case
		// k := RemoveDuplicates(numsInput)
		// k := RemoveDuplicatesV2(numsInput)
		k := RemoveDuplicatesV3(numsInput)

		// Slice the modified array to get the first 'k' elements
		actualArr := numsInput[:k]

		// Verify both the count and the content of the slice prefix
		passed := (k == tc.expectedK) && equal(actualArr, tc.expectedArr)

		status := "FAILED"
		if passed {
			passedCount++
			status = "PASSED"
		} else {
			failedCount++
			t.Errorf("Test Case %d failed", i+1) // Signals native 'go test' framework failure
		}

		fmt.Printf("Test Case %d: %s\n", i+1, status)
		fmt.Printf("  Input:    %v\n", tc.input)
		fmt.Printf("  Result k: %d (Expected: %d)\n", k, tc.expectedK)
		fmt.Printf("  Slice:    %v (Expected: %v)\n", actualArr, tc.expectedArr)
		fmt.Println("--------------------------------------------------")
	}

	// Final Summary Report
	fmt.Println("\n================ TEST SUMMARY ================")
	fmt.Printf("Total Test Cases: %d\n", len(testCases))
	fmt.Printf("Passed:           %d\n", passedCount)
	fmt.Printf("Failed:           %d\n", failedCount)
	fmt.Println("==============================================")
}
