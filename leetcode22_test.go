package main

import (
	"sort"
	"testing"
)

// Helper function to check if two string slices are equal regardless of element order
func equalSlices(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	// Copy slices to avoid mutating the originals
	sortedA := make([]string, len(a))
	sortedB := make([]string, len(b))
	copy(sortedA, a)
	copy(sortedB, b)

	sort.Strings(sortedA)
	sort.Strings(sortedB)

	for i := range sortedA {
		if sortedA[i] != sortedB[i] {
			return false
		}
	}
	return true
}

func TestGenerateParenthesis(t *testing.T) {
	// Define test cases (fixing the typo and count from before!)
	testCases := []struct {
		n        int
		expected []string
	}{
		{
			n:        1,
			expected: []string{"()"},
		},
		{
			n:        2,
			expected: []string{"(())", "()()"},
		},
		{
			n:        3,
			expected: []string{"((()))", "(()())", "(())()", "()(())", "()()()"},
		},
		{
			n:        0,
			expected: []string{""},
		},
		{
			n: 4,
			expected: []string{
				"(((())))", "((()()))", "((())())", "((()))()", "(()(()))",
				"(()()())", "(()())()", "(())(())", "(())()()", "()((()))",
				"()(()())", "()(())()", "()()(())", "()()()()",
			},
		},
	}

	successCount := 0
	failedCount := 0

	t.Log("Running Tests...")
	t.Log("--------------------------------------------------")

	for i, tc := range testCases {
		// actual := GenerateParenthesis(tc.n)
		// actual := GenerateParenthesisV2(tc.n)
		actual := GenerateParenthesisV3(tc.n)

		if equalSlices(actual, tc.expected) {
			t.Logf("Test Case %d (n=%d): PASSED", i+1, tc.n)
			successCount++
		} else {
			t.Errorf("Test Case %d (n=%d): FAILED\n  Expected: %v\n  Got:      %v", i+1, tc.n, tc.expected, actual)
			failedCount++
		}
	}

	t.Log("--------------------------------------------------")
	t.Logf("Total Successes: %d", successCount)
	t.Logf("Total Failures:  %d", failedCount)
}
