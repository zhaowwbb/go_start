package main

import (
	"fmt"
	"testing"
)

type TestCase struct {
	name     string
	digits   string
	expected []string
}

// Helper function to check if two string slices contain identical items regardless of order
func matchSlicesWithoutOrder(actual, expected []string) bool {
	if len(actual) != len(expected) {
		return false
	}

	// Transform the actual values into a frequencies map
	counts := make(map[string]int)
	for _, val := range actual {
		counts[val]++
	}

	// Verify all items in expected exist in the frequencies map
	for _, val := range expected {
		if counts[val] == 0 {
			return false
		}
		counts[val]--
	}
	return true
}

func TestLetterCombinationsSuite(t *testing.T) {
	testCases := []TestCase{
		{
			name:     "Standard LeetCode Example",
			digits:   "23",
			expected: []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"},
		},
		{
			name:     "Empty Input String",
			digits:   "",
			expected: []string{},
		},
		{
			name:     "Single Digit Input",
			digits:   "2",
			expected: []string{"a", "b", "c"},
		},
		{
			name:     "Digits with 4 Letters (7 and 9)",
			digits:   "7",
			expected: []string{"p", "q", "r", "s"},
		},
		{
			name:   "Longer Combination Mix",
			digits: "234",
			expected: []string{
				"adg", "adh", "adi", "aeg", "aeh", "aei", "afg", "afh", "afi",
				"bdg", "bdh", "bdi", "beg", "beh", "bei", "bfg", "bfh", "bfi",
				"cdg", "cdh", "cdi", "ceg", "ceh", "cei", "cfg", "cfh", "cfi",
			},
		},
		// Intentionally adding a failing test case to demonstrate failure counting logic
		// {
		// 	name:     "Deliberate Failure Case",
		// 	digits:   "2",
		// 	expected: []string{"x", "y", "z"},
		// },
	}

	passedCount := 0
	failedCount := 0

	fmt.Println("======================================================================")
	fmt.Printf("%48s\n", "RUNNING LETTER COMBINATIONS TEST SUITE")
	fmt.Println("======================================================================")

	for i, tc := range testCases {
		// The implementation is called exactly ONCE per test case here
		// actual := LetterCombinations(tc.digits)
		// actual := LetterCombinationsV2(tc.digits)
		actual := LetterCombinationsV3(tc.digits)

		if matchSlicesWithoutOrder(actual, tc.expected) {
			passedCount++
			fmt.Printf("Test %d: %-38s | ✅ PASSED (Generated %d items)\n", i+1, tc.name, len(actual))
		} else {
			failedCount++
			fmt.Printf("Test %d: %-38s | ❌ FAILED (Expected %d items, got %d)\n", i+1, tc.name, len(tc.expected), len(actual))
			t.Errorf("%s failed", tc.name)
		}
	}

	successRate := (float64(passedCount) / float64(len(testCases))) * 100

	fmt.Println("----------------------------------------------------------------------")
	fmt.Printf("TOTAL PASSED: %d\n", passedCount)
	fmt.Printf("TOTAL FAILED: %d\n", failedCount)
	fmt.Printf("SUCCESS RATE: %.1f%%\n", successRate)
	fmt.Println("======================================================================")
}
