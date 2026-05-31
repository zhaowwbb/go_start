// Package leetcode3 contains the solution for LeetCode 3: Longest Substring Without Repeating Characters.
package main

// LengthOfLongestSubstring finds the length of the longest substring without duplicate characters.
// Time Complexity: O(n) where n is the length of the string.
// Space Complexity: O(min(m, n)) where m is the character set size.
func LengthOfLongestSubstring(s string) int {
	charMap := make(map[rune]int)
	maxLength := 0
	left := 0

	// Convert string to runes to properly handle potential multi-byte UTF-8 characters
	for right, currentChar := range s {
		// If the character is a duplicate and exists within the current window
		if lastSeenIndex, exists := charMap[currentChar]; exists && lastSeenIndex >= left {
			// Slide the left pointer to the right of the duplicate character's position
			left = lastSeenIndex + 1
		}

		// Update or insert the current character's index position
		charMap[currentChar] = right

		// Calculate current window size and track maximum window size
		currentLength := right - left + 1
		if currentLength > maxLength {
			maxLength = currentLength
		}
	}

	return maxLength
}

func LengthOfLongestSubstringV2(s string) int {
	if len(s) == 0 {
		return 0
	}
	maxLen := 0
	left := 0
	charToIndex := make(map[rune]int)

	for right, c := range s {
		// c := s[right]
		if index, exists := charToIndex[c]; exists {
			if left <= index {
				left = index + 1
			}
		}
		charToIndex[c] = right
		currentLen := right - left + 1
		if currentLen > maxLen {
			maxLen = currentLen
		}
	}
	return maxLen
}

func LengthOfLongestSubstringV3(s string) int {
	charToIndex := make(map[rune]int)
	left := 0
	maxLen := 0

	for right, c := range s {
		if index, exists := charToIndex[c]; exists {
			if left <= index {
				left = index + 1
			}
		}
		charToIndex[c] = right

		currentLen := right - left + 1
		if currentLen > maxLen {
			maxLen = currentLen
		}
	}
	return maxLen
}
