package main

// LongestPalindrome finds the longest palindromic substring in s.
// Approach: Expand Around Center
// Time Complexity: O(n^2), Space Complexity: O(1)
func LongestPalindrome(s string) string {
	if len(s) < 1 {
		return ""
	}

	start, end := 0, 0

	for i := 0; i < len(s); i++ {
		// Case 1: Odd length palindrome (centered at i)
		len1 := expandAroundCenter(s, i, i)
		// Case 2: Even length palindrome (centered between i and i+1)
		len2 := expandAroundCenter(s, i, i+1)

		// Get the maximum length found at this center
		maxLen := len1
		if len2 > len1 {
			maxLen = len2
		}

		// If we found a longer palindrome, update our tracking indices
		if maxLen > (end - start) {
			start = i - (maxLen-1)/2
			end = i + maxLen/2
		}
	}

	// Slicing in Go is exclusive of the end index, so we use end + 1
	return s[start : end+1]
}

func expandAroundCenter(s string, left, right int) int {
	// Expand outward as long as characters match and indices are inbound
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
	}
	// Return the length of the palindrome found
	return right - left - 1
}

func LongestPalindromeV2(s string) string {
	left, right, maxLen := 0, 0, 0

	for i := range len(s) {
		len1 := expandFromCenterV2(s, i, i)
		len2 := expandFromCenterV2(s, i, i+1)
		if len1 > maxLen {
			maxLen = len1
		}
		if len2 > maxLen {
			maxLen = len2
		}
		if maxLen > right-left+1 {
			left = i - (maxLen-1)/2
			right = i + maxLen/2
		}
	}

	return s[left : right+1]
}

func expandFromCenterV2(s string, left int, right int) int {
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
	}
	return right - left - 1
}

func LongestPalindromeV3(s string) string {
	left, right, maxLen := 0, 0, 0
	for i := 0; i < len(s); i++ {
		len1 := expandFromCenterV3(s, i, i)
		len2 := expandFromCenterV3(s, i, i+1)
		maxLen = len1
		if maxLen < len2 {
			maxLen = len2
		}

		if maxLen > right-left+1 {
			left = i - (maxLen-1)/2
			right = i + maxLen/2
		}
	}
	return s[left : right+1]
}

func expandFromCenterV3(s string, left int, right int) int {
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
	}
	return right - left - 1
}
