package main

import "math"

// MyAtoi simulates the behavior of C/C++ atoi function.
// It parses a string to a 32-bit signed integer.
// Approach: Character scan with status tracking and overflow lookahead
// Time Complexity: O(n), Space Complexity: O(1)
func MyAtoi(s string) int {
	n := len(s)
	if n == 0 {
		return 0
	}

	i := 0

	// 1. Skip leading whitespaces
	for i < n && s[i] == ' ' {
		i++
	}

	if i == n {
		return 0
	}

	// 2. Read optional sign symbol
	sign := 1
	if s[i] == '+' {
		i++
	} else if s[i] == '-' {
		sign = -1
		i++
	}

	// 3. Extract and convert contiguous digit characters
	result := 0
	for i < n && s[i] >= '0' && s[i] <= '9' {
		digit := int(s[i] - '0')

		// Boundary Lookahead: Check for 32-bit signed integer limits before appending
		if sign == 1 {
			if result > math.MaxInt32/10 || (result == math.MaxInt32/10 && digit > 7) {
				return math.MaxInt32
			}
		} else {
			// Absolute threshold comparison for negative numbers
			if result > math.MaxInt32/10 || (result == math.MaxInt32/10 && digit > 8) {
				return math.MinInt32
			}
		}

		result = result*10 + digit
		i++
	}

	return sign * result
}

func MyAtoiV2(s string) int {
	if len(s) == 0 {
		return 0
	}
	n := len(s)
	i := 0
	for i < n && s[i] == ' ' {
		i++
	}
	if i == n {
		return 0
	}
	sign := 1
	if s[i] == '-' {
		sign = -1
		i++
	} else if s[i] == '+' {
		i++
	}

	result := 0
	for i < n && s[i] >= '0' && s[i] <= '9' {
		digit := int(s[i] - '0')

		if result > math.MaxInt32/10 || (result == math.MaxInt32/10 && digit > 7) {
			if sign > 0 {
				return math.MaxInt32
			} else {
				return math.MinInt32
			}
		}
		result = result*10 + digit
		i++
	}
	return result * sign
}

func MyAtoiV3(s string) int {
	n := len(s)
	if n == 0 {
		return 0
	}
	i := 0
	for i < n && s[i] == ' ' {
		i++
	}
	if i == n {
		return 0
	}
	sign := 1
	if s[i] == '-' {
		sign = -1
		i++
	}
	if s[i] == '+' {
		i++
	}
	result := 0
	for i < n && s[i] <= '9' && s[i] >= '0' {
		digit := int(s[i] - '0')
		if result > math.MaxInt32/10 || (result == math.MaxInt32/10 && digit > 7) {
			if sign > 0 {
				return math.MaxInt32
			} else {
				return math.MinInt32
			}
		}

		result = result*10 + digit
		i++
	}

	return result * sign
}
