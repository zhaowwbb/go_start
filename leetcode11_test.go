package main

import (
	"testing"
)

type testCase struct {
	height   []int
	expected int
}

func TestMaxArea(t *testing.T) {
	testCases := []testCase{
		{
			height:   []int{1, 8, 6, 2, 5, 4, 8, 3, 7},
			expected: 49,
		},
		{
			height:   []int{1, 1},
			expected: 1,
		},
		{
			height:   []int{4, 3, 2, 1, 4},
			expected: 16,
		},
		{
			height:   []int{1, 2, 1},
			expected: 2,
		},
	}

	for i, tc := range testCases {
		// SINGLE call execution invocation pattern
		// actual := MaxArea(tc.height)
		actual := MaxAreaV2(tc.height)

		if actual != tc.expected {
			t.Errorf("Test Case %d FAIL: Heights=%v\n Expected: %d\n Actual:   %d",
				i+1, tc.height, tc.expected, actual)
		} else {
			t.Logf("Test Case %d PASS: Heights=%v, Max Water=%d", i+1, tc.height, actual)
		}
	}
}
