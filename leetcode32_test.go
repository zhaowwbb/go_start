// leetcode32_test.go
package main

import (
	"fmt"
	"testing"
)

func TestLongestValidParentheses(t *testing.T) {
	// Table-driven test design (Go's idiomatic way to handle multiple inputs)
	testCases := []struct {
		input    string
		expected int
	}{
		{"(()", 2},
		{")()())", 4},
		{"", 0},
		{"()()()", 6},
		{"(())", 4},
		{"()(()", 2},
		{"(()))())(", 4},
		{"(((((", 0},
		{"))))", 0},
	}

	for _, tc := range testCases {
		// Call the function exactly once per case
		// actual := LongestValidParentheses(tc.input)
		actual := LongestValidParenthesesV2(tc.input)

		// pass := reflect.DeepEqual(nums, tc.expected)
		fmt.Printf("Input: %v Expected: %v Actual: %v \n",
			tc.input, tc.expected, actual)

		if actual != tc.expected {
			t.Errorf("FAIL\n  Input:    %q\n  Expected: %d\n  Actual:   %d\n", tc.input, tc.expected, actual)
		} else {
			t.Logf("PASS | Input: %q | Expected: %d | Actual: %d", tc.input, tc.expected, actual)
		}
	}
}
