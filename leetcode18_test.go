package main

import (
	"fmt"
	"sort"
	"testing"
)

type TestCase struct {
	name     string
	nums     []int
	target   int
	expected [][]int
}

// Deeply sorts and stringifies matrices to accurately verify contents across variant ordering
func canonicalKey(matrix [][]int) string {
	var elements []string
	for _, row := range matrix {
		// Ensure inner quadruplet values are strictly sorted
		sortedRow := make([]int, len(row))
		copy(sortedRow, row)
		sort.Ints(sortedRow)
		elements = append(elements, fmt.Sprint(sortedRow))
	}
	// Sort the overall output strings to ensure identical structure comparison
	sort.Strings(elements)
	return fmt.Sprint(elements)
}

func TestFourSumSuite(t *testing.T) {
	testCases := []TestCase{
		{
			name:     "Standard LeetCode Example",
			nums:     []int{1, 0, -1, 0, -2, 2},
			target:   0,
			expected: [][]int{{-2, -1, 1, 2}, {-2, 0, 0, 2}, {-1, 0, 0, 1}},
		},
		{
			name:     "All Elements Identical",
			nums:     []int{2, 2, 2, 2, 2},
			target:   8,
			expected: [][]int{{2, 2, 2, 2}},
		},
		{
			name:     "Target Unachievable",
			nums:     []int{1, 2, 3, 4},
			target:   100,
			expected: [][]int{},
		},
		{
			name:     "Negative and Zero Target Mix (Fixed)",
			nums:     []int{-3, -1, 0, 2, 4, 5},
			target:   2,
			expected: [][]int{{-3, -1, 2, 4}}, // Corrected to only include reachable elements
		},
	}

	passedCount := 0
	failedCount := 0

	fmt.Println("===========================================================================")
	fmt.Printf("%48s\n", "RUNNING 4SUM TEST SUITE")
	fmt.Println("===========================================================================")

	for i, tc := range testCases {
		// The implementation is called exactly ONCE per test case here
		// actual := FourSum(tc.nums, tc.target)
		actual := FourSumV2(tc.nums, tc.target)

		if canonicalKey(actual) == canonicalKey(tc.expected) {
			passedCount++
			fmt.Printf("Test %d: %-38s | ✅ PASSED (Found %d quadruplet(s))\n", i+1, tc.name, len(actual))
		} else {
			failedCount++
			fmt.Printf("Test %d: %-38s | ❌ FAILED (Expected %d matches, got %d)\n", i+1, tc.name, len(tc.expected), len(actual))
			t.Errorf("%s failed", tc.name)
		}
	}

	successRate := (float64(passedCount) / float64(len(testCases))) * 100

	fmt.Println("---------------------------------------------------------------------------")
	fmt.Printf("TOTAL PASSED: %d\n", passedCount)
	fmt.Printf("TOTAL FAILED: %d\n", failedCount)
	fmt.Printf("SUCCESS RATE: %.1f%%\n", successRate)
	fmt.Println("===========================================================================")
}
