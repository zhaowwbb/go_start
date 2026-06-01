package main

import (
	"testing"
)

func TestLongestPalindrome(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected []string // Using a slice because multiple answers can be valid (e.g., "bab" or "aba")
	}{
		{
			name:     "Standard mixed string",
			input:    "babad",
			expected: []string{"bab", "aba"},
		},
		{
			name:     "Even-length palindrome",
			input:    "cbbd",
			expected: []string{"bb"},
		},
		{
			name:     "Single character",
			input:    "a",
			expected: []string{"a"},
		},
		{
			name:     "Two distinct characters",
			input:    "ac",
			expected: []string{"a", "c"},
		},
		{
			name:     "Full string palindrome",
			input:    "racecar",
			expected: []string{"racecar"},
		},
		{
			name:     "Longer string with internal palindrome",
			input:    "aacabdkacaa",
			expected: []string{"aca"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// actual := LongestPalindrome(tt.input)
			// actual := LongestPalindromeV2(tt.input)
			actual := LongestPalindromeV3(tt.input)

			// Verify if the result matches any of the valid expected options
			matched := false
			for _, exp := range tt.expected {
				if actual == exp {
					matched = true
					break
				}
			}

			if !matched {
				t.Errorf("LongestPalindrome(%q) = %q; want one of %v", tt.input, actual, tt.expected)
			}
		})
	}
}
