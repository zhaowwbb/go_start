package main

import (
	"fmt"
	"sort"
	"testing"
)

// TestCase defines the structure for our test data
type TestCase struct {
	input       []int
	val         int
	expectedK   int
	expectedArr []int // The first k elements (sorted for order-agnostic comparison)
}

// verifySlice checks if the first k elements match the expected elements
func verifySlice(actual []int, expected []int) bool {
	if len(actual) != len(expected) {
		return false
	}

	// Create copies to sort because LeetCode accepts elements in any order
	aCopy := make([]int, len(actual))
	eCopy := make([]int, len(expected))
	copy(aCopy, actual)
	copy(eCopy, expected)

	sort.Ints(aCopy)
	sort.Ints(eCopy)

	for i := range aCopy {
		if aCopy[i] != eCopy[i] {
			return false
		}
	}
	return true
}

func TestRemoveElement(t *testing.T) {
	testCases := []TestCase{
		{input: []int{3, 2, 2, 3}, val: 3, expectedK: 2, expectedArr: []int{2, 2}},
		{input: []int{0, 1, 2, 2, 3, 0, 4, 2}, val: 2, expectedK: 5, expectedArr: []int{0, 1, 3, 0, 4}},
		{input: []int{}, val: 1, expectedK: 0, expectedArr: []int{}},
		{input: []int{1}, val: 1, expectedK: 0, expectedArr: []int{}},
		{input: []int{1}, val: 2, expectedK: 1, expectedArr: []int{1}},
		{input: []int{2, 2, 2}, val: 2, expectedK: 0, expectedArr: []int{}},
	}

	fmt.Println("Running test cases for Remove Element...")
	fmt.Println("--------------------------------------------------")

	passedCount := 0
	failedCount := 0

	for i, tc := range testCases {
		// Create a deep copy of input slice to preserve original for debugging output
		numsInput := make([]int, len(tc.input))
		copy(numsInput, tc.input)

		// Call the implementation exactly once per test case
		// k := RemoveElement(numsInput, tc.val)
		// k := RemoveElementV2(numsInput, tc.val)
		k := RemoveElementV3(numsInput, tc.val)

		// Extract the first k elements modified by the function
		actualArr := numsInput[:k]

		// Verify the length count and the correct elements remain
		passed := (k == tc.expectedK) && verifySlice(actualArr, tc.expectedArr)

		status := "FAILED"
		if passed {
			passedCount++
			status = "PASSED"
		} else {
			failedCount++
			t.Errorf("Test Case %d failed", i+1)
		}

		fmt.Printf("Test Case %d: %s\n", i+1, status)
		fmt.Printf("  Input Array: %v, Target Value: %d\n", tc.input, tc.val)
		fmt.Printf("  Result k:    %d (Expected: %d)\n", k, tc.expectedK)
		fmt.Printf("  Slice Prefix:%v (Expected elements: %v)\n", actualArr, tc.expectedArr)
		fmt.Println("--------------------------------------------------")
	}

	// Final Summary Report
	fmt.Println("\n================ TEST SUMMARY ================")
	fmt.Printf("Total Test Cases: %d\n", len(testCases))
	fmt.Printf("Passed:           %d\n", passedCount)
	fmt.Printf("Failed:           %d\n", failedCount)
	fmt.Println("==============================================")
}
