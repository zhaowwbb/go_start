package main

import (
	"fmt"
	"testing"
)

func TestIsMatch(t *testing.T) {
	// Table-driven test cases array mapping to the requested pattern
	testCases := []struct {
		s        string
		p        string
		expected bool
		desc     string
	}{
		{s: "aa", p: "a", expected: false, desc: "Exact character mismatch (length difference)"},
		{s: "aa", p: "a*", expected: true, desc: "Star operator (*) matching multiple characters"},
		{s: "ab", p: ".*", expected: true, desc: "Dot-Star (.*) matching any sequence"},
		{s: "aab", p: "c*a*b", expected: true, desc: "Preceding element with 0-occurrence star matching"},
		{s: "mississippi", p: "mis*is*p*.", expected: false, desc: "Complex combination resulting in mismatch"},
		{s: "ab", p: ".*c", expected: false, desc: "Dot-Star match fails due to trailing character requirement"},
	}

	fmt.Println("Running LeetCode 10 - Regular Expression Matching Tests:")
	fmt.Println("=================================================================")

	passed := 0

	for i, tc := range testCases {
		// Call the function exactly ONE time per test case execution
		// result := IsMatch(tc.s, tc.p)
		// result := IsMatchV2(tc.s, tc.p)
		result := IsMatchV3(tc.s, tc.p)

		fmt.Printf("Test %d: %s\n", i+1, tc.desc)
		fmt.Printf("  Input String  (s): %s\n", tc.s)
		fmt.Printf("  Input Pattern (p): %s\n", tc.p)
		fmt.Printf("  Expected Result:   %t\n", tc.expected)
		fmt.Printf("  Actual Result:     %t\n", result)

		if result == tc.expected {
			fmt.Println("  Status:            PASSED ✅")
			passed++
		} else {
			fmt.Println("  Status:            FAILED ❌")
			t.Errorf("Test %d Failed (%s): Input s=%q, p=%q expected %t, got %t", i+1, tc.desc, tc.s, tc.p, tc.expected, result)
		}
		fmt.Println("---------------------------------------------")
	}

	fmt.Printf("\nResult: %d/%d tests passed.\n", passed, len(testCases))
}
