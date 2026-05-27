package main

import (
	"fmt"
	"testing"
)

func TestSearchInRotatedSortedArray(t *testing.T) {
	// Define the table-driven test cases structure
	type testCase struct {
		nums     []int
		target   int
		expected int
		desc     string
	}

	testCases := []testCase{
		{[]int{4, 5, 6, 7, 0, 1, 2}, 0, 4, "Example 1: Target exists in the right rotated part"},
		{[]int{4, 5, 6, 7, 0, 1, 2}, 3, -1, "Example 2: Target does not exist"},
		{[]int{1}, 0, -1, "Example 3: Single element array, target does not exist"},
		{[]int{1}, 1, 0, "Single element array, target exists"},
		{[]int{5, 1, 3}, 3, 2, "Small rotated array, target on the right"},
		{[]int{3, 5, 1}, 3, 0, "Small rotated array, target on the left"},
	}

	fmt.Println("\nRunning LeetCode 33 - Search in Rotated Sorted Array Tests:")
	fmt.Println("============================================================")

	for i, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			// Each test case invokes the function exactly once
			// actual := searchInRotatedSortedArray(tc.nums, tc.target)
			// actual := searchInRotatedSortedArrayV2(tc.nums, tc.target)
			actual := searchInRotatedSortedArrayV3(tc.nums, tc.target)

			// Print output exactly as specified
			fmt.Printf("Test %d: %s\n", i+1, tc.desc)
			fmt.Printf("Input: nums=%v, target=%d Expected: %v Actual: %v \n\n",
				tc.nums, tc.target, tc.expected, actual)

			if actual != tc.expected {
				t.Errorf("Test failed! Expected %d, but got %d", tc.expected, actual)
			}
		})
	}
	fmt.Println("============================================================")
}
