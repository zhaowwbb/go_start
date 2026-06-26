package main

import (
	"fmt"
	"testing"
)

// TestCase defines the structure for our test inputs and expectations
type TestCase struct {
	inputs   []string
	expected string
}

func TestLongestCommonPrefix(t *testing.T) {
	testCases := []TestCase{
		{[]string{"flower", "flow", "flight"}, "fl"},
		{[]string{"dog", "racecar", "car"}, ""},
		{[]string{"interstate", "interview", "internal"}, "inter"},
		{[]string{"alone"}, "alone"},
		{[]string{"", "b", "c"}, ""},
		{[]string{"test", "test", "test"}, "test"},
		{[]string{}, ""},
	}

	passedCount := 0
	failedCount := 0

	fmt.Println("--- Running LeetCode 14 Tests (Go) ---")

	for i, tc := range testCases {
		caseNum := i + 1

		// Call the implementation exactly once
		// result := LongestCommonPrefix(tc.inputs)
		// result := LongestCommonPrefixV2(tc.inputs)
		result := LongestCommonPrefixV3(tc.inputs)

		if result == tc.expected {
			fmt.Printf("Test Case %d: PASSED\n", caseNum)
			passedCount++
		} else {
			fmt.Printf("Test Case %d: FAILED (Expected '%s', got '%s')\n", caseNum, tc.expected, result)
			failedCount++
			t.Errorf("Case %d failed", caseNum) // Marks the overall go test as failed if a case drops
		}
	}

	fmt.Println("---------------------------------")
	fmt.Printf("Summary: %d passed, %d failed.\n", passedCount, failedCount)
}
