package main

// LongestCommonPrefix finds the longest common prefix string amongst an array of strings.
func LongestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	// Iterate through the characters of the first string
	for i := 0; i < len(strs[0]); i++ {
		char := strs[0][i]

		// Check this character against the rest of the strings
		for _, str := range strs[1:] {
			// If the current string is shorter than i, or characters don't match
			if i == len(str) || str[i] != char {
				return strs[0][:i]
			}
		}
	}

	return strs[0]
}

func LongestCommonPrefixV2(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	first := strs[0]
	for i := 0; i < len(first); i++ {
		currentChar := first[i]
		for j := 1; j < len(strs); j++ {
			s := strs[j]
			if i == len(s) || currentChar != s[i] {
				return first[:i]
			}
		}
	}
	return first
}

func LongestCommonPrefixV3(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	for i := 0; i < len(strs[0]); i++ {
		current := strs[0][i]
		for _, str := range strs[1:] {
			if i == len(str) || current != str[i] {
				return strs[0][:i]
			}
		}
	}

	return strs[0]
}
