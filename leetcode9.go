package main

// IsPalindrome determines if an integer x is a palindrome without converting it to a string.
func IsPalindrome(x int) bool {
	// Special cases:
	// 1. Negative numbers are not palindromes (e.g., -121)
	// 2. If the last digit is 0, the first digit must also be 0 (only '0' satisfies this)
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}

	reversedNumber := 0
	// Reversing the second half of the number
	for x > reversedNumber {
		pop := x % 10
		reversedNumber = reversedNumber*10 + pop
		x /= 10
	}

	// When the length is an odd number, we can get rid of the middle digit by reversedNumber / 10
	return x == reversedNumber || x == reversedNumber/10
}

func IsPalindromeV2(x int) bool {
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}

	reverse := 0
	for x > reverse {
		pop := x % 10
		reverse = reverse*10 + pop
		x /= 10
	}
	return x == reverse || x == reverse/10

}
