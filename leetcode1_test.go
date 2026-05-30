package main

import (
	"fmt"
	"reflect"
	"testing"
)

// TwoSumTestCase defines the structure for our test group
type TwoSumTestCase struct {
	id       int
	nums     []int
	target   int
	expected []int
}

func TestTwoSum(t *testing.T) {
	// Test case group containing all scenario data
	testCases := []TwoSumTestCase{
		{
			id:       1,
			nums:     []int{2, 7, 11, 15},
			target:   9,
			expected: []int{0, 1},
		},
		{
			id:       2,
			nums:     []int{3, 2, 4},
			target:   6,
			expected: []int{1, 2},
		},
		{
			id:       3,
			nums:     []int{3, 3},
			target:   6,
			expected: []int{0, 1},
		},
	}

	// Loop through the test case group
	for _, tc := range testCases {
		// Call implementation function exactly once
		// actualResult := twoSum(tc.nums, tc.target)
		// actualResult := twoSumV2(tc.nums, tc.target)
		actualResult := twoSumV3(tc.nums, tc.target)

		fmt.Printf("Test Case %d:\n", tc.id)
		fmt.Printf("  nums:     %v\n", tc.nums)
		fmt.Printf("  target:   %d\n", tc.target)
		fmt.Printf("  Expected: %v\n", tc.expected)
		fmt.Printf("  Actual:   %v\n", actualResult)
		fmt.Println("----------------------------------------")

		// Verify correctness
		if !reflect.DeepEqual(actualResult, tc.expected) {
			t.Errorf("Test Case %d failed. Expected %v, got %v", tc.id, tc.expected, actualResult)
		}
	}
}
