package main

import (
	"fmt"
	"reflect"
	"testing"
)

// TestCase defines the structure for our test group
type TestCase struct {
	id       int
	s        string
	words    []string
	expected []int
}

func TestFindSubstring(t *testing.T) {
	// Test case group containing all scenario data
	testCases := []TestCase{
		{
			id:       1,
			s:        "barfoothefoobarman",
			words:    []string{"foo", "bar"},
			expected: []int{0, 9},
		},
		{
			id:       2,
			s:        "wordgoodgoodgoodbestword",
			words:    []string{"word", "good", "best", "word"},
			expected: []int{},
		},
		{
			id:       3,
			s:        "barfoofoobarthefoobarman",
			words:    []string{"bar", "foo", "the"},
			expected: []int{6, 9, 12},
		},
		{
			id:       4,
			s:        "wordgoodgoodgoodbestword",
			words:    []string{"word", "good", "best", "good"},
			expected: []int{8},
		},
	}

	// Loop through the test case group
	for _, tc := range testCases {
		// Call implementation function exactly once
		// actualResult := findSubstring(tc.s, tc.words)
		// actualResult := findSubstringV2(tc.s, tc.words)
		actualResult := findSubstringV3(tc.s, tc.words)

		fmt.Printf("Test Case %d:\n", tc.id)
		fmt.Printf("  s:        '%s'\n", tc.s)
		fmt.Printf("  words:    %v\n", tc.words)
		fmt.Printf("  Expected: %v\n", tc.expected)
		fmt.Printf("  Actual:   %v\n", actualResult)
		fmt.Println("----------------------------------------")

		// Verify correctness (handles empty slice vs nil slice variants natively)
		if len(actualResult) == 0 && len(tc.expected) == 0 {
			continue
		}
		if !reflect.DeepEqual(actualResult, tc.expected) {
			t.Errorf("Test Case %d failed. Expected %v, got %v", tc.id, tc.expected, actualResult)
		}
	}
}
