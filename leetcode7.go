package main

import "math"

// Reverse takes a signed 32-bit integer x and returns its digits reversed.
// If reversing x causes the value to overflow/underflow a 32-bit integer range, it returns 0.
// Approach: Digit extraction with strict 32-bit boundary checks.
// Time Complexity: O(log10(x)) -> effectively O(1), Space Complexity: O(1)
func Reverse(x int) int {
	reversed := 0

	for x != 0 {
		// Extract the last digit (Go's % operator handles negative numbers correctly)
		pop := x % 10
		x /= 10

		// Check for 32-bit signed integer overflow before multiplying by 10
		// math.MaxInt32 is 2147483647
		if reversed > math.MaxInt32/10 || (reversed == math.MaxInt32/10 && pop > 7) {
			return 0
		}

		// Check for 32-bit signed integer underflow before multiplying by 10
		// math.MinInt32 is -2147483648
		if reversed < math.MinInt32/10 || (reversed == math.MinInt32/10 && pop < -8) {
			return 0
		}

		// Append the digit securely
		reversed = reversed*10 + pop
	}

	return reversed
}

func ReverseV2(x int) int {
	rev := 0

	for x != 0 {
		pop := x % 10
		x /= 10
		if rev > math.MaxInt32/10 || (rev == math.MaxInt32/10 && pop > 7) {
			return 0
		}
		if rev < math.MinInt32/10 || (rev == math.MaxInt32/10 && pop < -8) {
			return 0
		}
		rev = rev*10 + pop
	}
	return rev
}

func ReverseV3(x int) int {
	rev := 0
	for x != 0 {
		pop := x % 10
		x /= 10
		if rev > math.MaxInt32/10 || (rev == math.MaxInt32/10 && pop > 7) {
			return 0
		}
		if rev < math.MinInt32/10 || (rev == math.MinInt32 && pop < -8) {
			return 0
		}
		rev = rev*10 + pop
	}
	return rev
}
