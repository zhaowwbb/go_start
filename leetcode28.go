package main

// StrStr finds the index of the first occurrence of needle in haystack.
// Returns -1 if needle is not part of haystack.
func StrStr(haystack string, needle string) int {
	// Edge case: An empty needle is always found at index 0
	if needle == "" {
		return 0
	}

	n := len(haystack)
	m := len(needle)

	// Only iterate up to the point where the remaining haystack is at least as long as needle
	for i := 0; i <= n-m; i++ {
		// Check if the substring from index i to i+m matches the needle
		if haystack[i:i+m] == needle {
			return i
		}
	}

	return -1
}

func StrStrV2(haystack string, needle string) int {
	if needle == "" {
		return 0
	}
	n := len(haystack)
	m := len(needle)
	for i := range n - m + 1 {
		if haystack[i:i+m] == needle {
			return i
		}
	}
	return -1
}

func StrStrV3(haystack string, needle string) int {
	if needle == "" {
		return 0
	}
	n, m := len(haystack), len(needle)
	for i := 0; i <= n-m; i++ {
		if haystack[i:i+m] == needle {
			return i
		}
	}
	return -1
}
