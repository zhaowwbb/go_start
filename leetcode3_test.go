package main

import (
	"testing"
)

func TestLengthOfLongestSubstring(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected int
	}{
		{"Standard mixed repeating characters", "abcabcbb", 3},
		{"All identical characters", "bbbbb", 1},
		{"Substring inside a subsequence pattern", "pwwkew", 3},
		{"Empty string boundary case", "", 0},
		{"Single space string", " ", 1},
		{"Two distinct characters", "au", 2},
		{"Repeated character separated by others", "dvdf", 3},
		{"All unique characters", "abcdefghijklmnopqrstuvwxyz", 26},
		{"Includes digits, symbols, and spaces", "1234!@#$ 1234", 9},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// actual := LengthOfLongestSubstring(tc.input)
			// actual := LengthOfLongestSubstringV2(tc.input)
			actual := LengthOfLongestSubstringV3(tc.input)
			if actual != tc.expected {
				t.Errorf("LengthOfLongestSubstring(%q) = %d; expected %d", tc.input, actual, tc.expected)
			}
		})
	}
}
