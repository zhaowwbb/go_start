package main

import (
	"reflect"
	"sort"
	"testing"
)

// Helper structural type to handle test matrix details
type testCase struct {
	candidates []int
	target     int
	expected   [][]int
}

// Helper function to deeply sort a 2D matrix to reliably test list equality
func normalizeMatrix(matrix [][]int) [][]int {
	if matrix == nil {
		return [][]int{}
	}

	// First sort each inner array
	for _, row := range matrix {
		sort.Ints(row)
	}

	// Then sort the outer matrix arrays by element-wise value
	sort.Slice(matrix, func(i, j int) bool {
		minLen := len(matrix[i])
		if len(matrix[j]) < minLen {
			minLen = len(matrix[j])
		}
		for k := 0; k < minLen; k++ {
			if matrix[i][k] != matrix[j][k] {
				return matrix[i][k] < matrix[j][k]
			}
		}
		return len(matrix[i]) < len(matrix[j])
	})
	return matrix
}

func TestCombinationSum(t *testing.T) {
	testCases := []testCase{
		{
			candidates: []int{2, 3, 6, 7},
			target:     7,
			expected:   [][]int{{2, 2, 3}, {7}},
		},
		{
			candidates: []int{2, 3, 5},
			target:     8,
			expected:   [][]int{{2, 2, 2, 2}, {2, 3, 3}, {3, 5}},
		},
		{
			candidates: []int{2},
			target:     1,
			expected:   [][]int{},
		},
		{
			candidates: []int{5, 10, 15},
			target:     5,
			expected:   [][]int{{5}},
		},
	}

	for i, tc := range testCases {
		// SINGLE call execution invocation pattern
		// actual := CombinationSum(tc.candidates, tc.target)
		// actual := CombinationSumV2(tc.candidates, tc.target)
		actual := CombinationSumV3(tc.candidates, tc.target)

		// Normalize sorting order of matrix data patterns for deterministic mapping
		sortedActual := normalizeMatrix(actual)
		sortedExpected := normalizeMatrix(tc.expected)

		if !reflect.DeepEqual(sortedActual, sortedExpected) {
			t.Errorf("Test Case %d FAIL: Candidates=%v, Target=%d\n Expected: %v\n Actual:   %v",
				i+1, tc.candidates, tc.target, sortedExpected, sortedActual)
		} else {
			t.Logf("Test Case %d PASS: Candidates=%v, Target=%d", i+1, tc.candidates, tc.target)
		}
	}
}
