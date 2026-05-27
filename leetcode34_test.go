package main

import (
	"fmt"
	"testing"
)

func TestSearchRange(t *testing.T) {
	// Define the table-driven test cases structure
	type testCase struct {
		nums     []int
		target   int
		expected []int
		desc     string
	}

	testCases := []testCase{
		{[]int{5, 7, 7, 8, 8, 10}, 8, []int{3, 4}, "Example 1: Target exists multiple times"},
		{[]int{5, 7, 7, 8, 8, 10}, 6, []int{-1, -1}, "Example 2: Target does not exist inside array bounds"},
		{[]int{}, 0, []int{-1, -1}, "Example 3: Empty array"},
		{[]int{2, 2}, 2, []int{0, 1}, "All elements match the target"},
		{[]int{1, 2, 3, 4, 5}, 3, []int{2, 2}, "Target appears exactly once in the middle"},
		{[]int{1, 3, 5, 7, 9}, 10, []int{-1, -1}, "Target is greater than all elements"},
	}

	fmt.Println("\nRunning LeetCode 34 - Find First and Last Position Tests:")
	fmt.Println("============================================================")

	for i, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			// Each test case invokes the function exactly once
			// actual := searchRange(tc.nums, tc.target)
			// actual := searchRangeV2(tc.nums, tc.target)
			actual := searchRangeV3(tc.nums, tc.target)

			// Print output tracking execution state
			fmt.Printf("Test %d: %s\n", i+1, tc.desc)
			fmt.Printf("Input: nums=%v, target=%d Expected: %v Actual: %v \n\n",
				tc.nums, tc.target, tc.expected, actual)

			// Slice comparison validation logic
			if actual[0] != tc.expected[0] || actual[1] != tc.expected[1] {
				t.Errorf("Test failed! Expected %v, but got %v", tc.expected, actual)
			}
		})
	}
	fmt.Println("============================================================")
}
