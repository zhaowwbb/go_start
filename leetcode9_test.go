package main

import (
	"fmt"
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	// Test case struct matching the requested configuration
	testCases := []struct {
		input    int
		expected bool
		desc     string
	}{
		{input: 121, expected: true, desc: "Standard odd-length palindrome"},
		{input: -121, expected: false, desc: "Negative number, not a palindrome"},
		{input: 10, expected: false, desc: "Ends with 0 but is not 0, not a palindrome"},
		{input: 0, expected: true, desc: "Single digit zero is a palindrome"},
		{input: 7, expected: true, desc: "Single digit non-zero is a palindrome"},
		{input: 123321, expected: true, desc: "Standard even-length palindrome"},
	}

	fmt.Println("Running LeetCode 9 - Palindrome Number Tests:")
	fmt.Println("============================================================")

	passed := 0

	for i, tc := range testCases {
		// Call the function exactly ONE time per test case
		// actual := IsPalindrome(tc.input)
		actual := IsPalindromeV2(tc.input)

		fmt.Printf("Test %d: %s\n", i+1, tc.desc)
		fmt.Printf("  Input:    %d\n", tc.input)
		fmt.Printf("  Expected: %t\n", tc.expected)
		fmt.Printf("  Actual:   %t\n", actual)

		if actual == tc.expected {
			fmt.Println("  Status:   PASSED ✅")
			passed++
		} else {
			fmt.Println("  Status:   FAILED ❌")
			t.Errorf("Test %d Failed (%s): Input %d expected %t, got %t", i+1, tc.desc, tc.input, tc.expected, actual)
		}
		fmt.Println("----------------------------------------")
	}

	fmt.Printf("\nResult: %d/%d tests passed.\n", passed, len(testCases))
}
