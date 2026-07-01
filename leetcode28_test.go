package main

import (
	"fmt"
	"testing"
)

// TestCase defines the structure for our string-matching tests
type TestCase struct {
	haystack      string
	needle        string
	expectedIndex int
}

func TestStrStr(t *testing.T) {
	testCases := []TestCase{
		{haystack: "sadbutsad", needle: "sad", expectedIndex: 0},
		{haystack: "leetcode", needle: "leeto", expectedIndex: -1},
		{haystack: "hello", needle: "ll", expectedIndex: 2},
		{haystack: "a", needle: "a", expectedIndex: 0},
		{haystack: "abc", needle: "", expectedIndex: 0},
		{haystack: "mississippi", needle: "issip", expectedIndex: 4},
	}

	fmt.Println("Running test cases for Find First Occurrence...")
	fmt.Println("--------------------------------------------------")

	passedCount := 0
	failedCount := 0

	for i, tc := range testCases {
		// Call the implementation exactly once per test case
		// result := StrStr(tc.haystack, tc.needle)
		// result := StrStrV2(tc.haystack, tc.needle)
		result := StrStrV3(tc.haystack, tc.needle)

		passed := result == tc.expectedIndex

		status := "FAILED"
		if passed {
			passedCount++
			status = "PASSED"
		} else {
			failedCount++
			t.Errorf("Test Case %d failed", i+1)
		}

		fmt.Printf("Test Case %d: %s\n", i+1, status)
		fmt.Printf("  Haystack:     '%s' | Needle: '%s'\n", tc.haystack, tc.needle)
		fmt.Printf("  Result Index: %d (Expected: %d)\n", result, tc.expectedIndex)
		fmt.Println("--------------------------------------------------")
	}

	// Final Summary Report
	fmt.Println("\n================ TEST SUMMARY ================")
	fmt.Printf("Total Test Cases: %d\n", len(testCases))
	fmt.Printf("Passed:           %d\n", passedCount)
	fmt.Printf("Failed:           %d\n", failedCount)
	fmt.Println("==============================================")
}
