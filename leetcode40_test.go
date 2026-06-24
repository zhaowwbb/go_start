package main

import (
	"reflect"
	"sort"
	"testing"
)

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

	for _, row := range matrix {
		sort.Ints(row)
	}

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

func TestCombinationSum2(t *testing.T) {
	testCases := []testCase{
		{
			candidates: []int{10, 1, 2, 7, 6, 1, 5},
			target:     8,
			expected:   [][]int{{1, 1, 6}, {1, 2, 5}, {1, 7}, {2, 6}},
		},
		{
			candidates: []int{2, 5, 2, 1, 2},
			target:     5,
			expected:   [][]int{{1, 2, 2}, {5}},
		},
		{
			candidates: []int{2, 4, 6},
			target:     5,
			expected:   [][]int{},
		},
	}

	for i, tc := range testCases {
		// SINGLE call execution per test case requirement
		// actual := CombinationSum2(tc.candidates, tc.target)
		actual := CombinationSum2V2(tc.candidates, tc.target)

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
