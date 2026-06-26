package main

import (
	"fmt"
	"reflect"
	"sort"
	"testing"
)

// TestCase defines the structure for our test inputs and expectations
type TestCase struct {
	inputs   []int
	expected [][]int
}

// Helper to deeply sort 2D slices so that verification is order-agnostic
func canonicalize(slices [][]int) [][]int {
	if slices == nil {
		return [][]int{}
	}

	// Sort elements inside each individual triplet first
	for _, slice := range slices {
		sort.Ints(slice)
	}

	// Sort the outer list of triplets
	sort.Slice(slices, func(i, j int) bool {
		if slices[i][0] != slices[j][0] {
			return slices[i][0] < slices[j][0]
		}
		if slices[i][1] != slices[j][1] {
			return slices[i][1] < slices[j][1]
		}
		return slices[i][2] < slices[j][2]
	})

	return slices
}

func TestThreeSum(t *testing.T) {
	testCases := []TestCase{
		{[]int{-1, 0, 1, 2, -1, -4}, [][]int{{-1, -1, 2}, {-1, 0, 1}}},
		{[]int{0, 1, 1}, [][]int{}},
		{[]int{0, 0, 0}, [][]int{{0, 0, 0}}},
		{[]int{-2, 0, 0, 2, 2}, [][]int{{-2, 0, 2}}},
		{[]int{-4, -2, -2, -2, 0, 1, 2, 2, 2, 3, 4}, [][]int{
			{-4, 0, 4}, {-4, 1, 3}, {-4, 2, 2}, {-2, -2, 4}, {-2, 0, 2},
		}},
		{[]int{}, [][]int{}},
		{[]int{1, 2}, [][]int{}},
	}

	passedCount := 0
	failedCount := 0

	fmt.Println("--- Running LeetCode 15 Tests (Go 3Sum) ---")

	for i, tc := range testCases {
		caseNum := i + 1

		// Call the implementation exactly once per test case
		// result := ThreeSum(tc.inputs)
		// result := ThreeSumV2(tc.inputs)
		result := ThreeSumV3(tc.inputs)

		// Normalize both expected and actual results to compare cleanly
		actualSorted := canonicalize(result)
		expectedSorted := canonicalize(tc.expected)

		if reflect.DeepEqual(actualSorted, expectedSorted) {
			fmt.Printf("Test Case %d: PASSED\n", caseNum)
			passedCount++
		} else {
			fmt.Printf("Test Case %d: FAILED\n", caseNum)
			fmt.Printf("   Expected: %v\n", tc.expected)
			fmt.Printf("   Got:      %v\n", result)
			failedCount++
			t.Errorf("Case %d failed", caseNum)
		}
	}

	fmt.Println("-------------------------------------------")
	fmt.Printf("Summary: %d passed, %d failed.\n", passedCount, failedCount)
}
