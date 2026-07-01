package main

import "math"

// Divide performs integer division of two integers without using
// multiplication, division, or the mod operator.
func Divide(dividend int, divisor int) int {
	// Handle overflow edge cases as per LeetCode constraints
	if dividend == math.MinInt32 && divisor == -1 {
		return math.MaxInt32
	}
	if dividend == math.MinInt32 && divisor == 1 {
		return math.MinInt32
	}

	// Determine the sign of the result
	// True if the signs are different
	isNegative := (dividend < 0) != (divisor < 0)

	// Work with absolute values using 64-bit integers to prevent overflow during negation
	absDividend := int64(dividend)
	if absDividend < 0 {
		absDividend = -absDividend
	}

	absDivisor := int64(divisor)
	if absDivisor < 0 {
		absDivisor = -absDivisor
	}

	var quotient int64 = 0

	// Perform bit-shifting division (similar to long division)
	for absDividend >= absDivisor {
		var tempDivisor = absDivisor
		var multiple int64 = 1

		// Double the divisor and the multiple as long as it fits inside the remaining dividend
		for absDividend >= (tempDivisor << 1) {
			tempDivisor <<= 1
			multiple <<= 1
		}

		// Reduce the dividend and accumulate the quotient
		absDividend -= tempDivisor
		quotient += multiple
	}

	// Apply the sign to the quotient
	if isNegative {
		quotient = -quotient
	}

	// Double-check integer bounds before returning
	if quotient > math.MaxInt32 {
		return math.MaxInt32
	}
	if quotient < math.MinInt32 {
		return math.MinInt32
	}

	return int(quotient)
}

func DivideV2(dividend int, divisor int) int {
	if dividend == math.MinInt32 && divisor == -1 {
		return math.MaxInt32
	}
	if dividend == math.MinInt32 && divisor == 1 {
		return math.MinInt32
	}

	sign := 1
	if dividend > 0 && divisor < 0 {
		sign = -1
	}
	if dividend < 0 && divisor > 0 {
		sign = -1
	}
	if dividend < 0 {
		dividend = 0 - dividend
	}
	if divisor < 0 {
		divisor = 0 - divisor
	}
	result := 0
	for dividend >= divisor {
		dividend -= divisor
		result++
	}
	return result * sign
}

func DivideV3(dividend int, divisor int) int {
	if dividend == math.MinInt32 && divisor == -1 {
		return math.MaxInt32
	}
	isNegative := (dividend < 0) != (divisor < 0)

	absDividend := int64(dividend)
	if absDividend < 0 {
		absDividend = 0 - absDividend
	}
	absDivisor := int64(divisor)
	var result int64 = 0
	if absDivisor < 0 {
		absDivisor = 0 - absDivisor
	}
	for absDividend >= absDivisor {
		var tempDivisor = absDivisor
		var multiplier int64 = 1
		for absDividend >= (tempDivisor << 1) {
			multiplier = multiplier << 1
			tempDivisor = tempDivisor << 1
		}
		result += multiplier
		absDividend -= tempDivisor
	}

	if isNegative {
		result = 0 - result
	}

	if result > math.MaxInt32 {
		return math.MaxInt32
	}
	if result < math.MinInt32 {
		return math.MinInt32
	}

	return int(result)
}

func DivideV4(dividend int, divisor int) int {
	if dividend == math.MinInt32 && divisor == -1 {
		return math.MaxInt32
	}
	if dividend == math.MinInt32 && divisor == 1 {
		return math.MinInt32
	}

	isNegative := (dividend < 0) != (divisor < 0)
	absDividend := int64(dividend)
	if absDividend < 0 {
		absDividend = 0 - absDividend
	}
	absDivisor := int64(divisor)
	if absDivisor < 0 {
		absDivisor = 0 - absDivisor
	}
	var quotient int64 = 0
	for absDividend >= absDivisor {
		var tempDivisor int64 = absDivisor
		var multiple int64 = 1
		for absDividend > tempDivisor<<1 {
			tempDivisor <<= 1
			multiple <<= 1
		}

		absDividend -= tempDivisor
		quotient += multiple
	}

	if quotient > math.MaxInt32 {
		return math.MaxInt32
	}
	if quotient < math.MinInt32 {
		return math.MinInt32
	}

	if isNegative {
		quotient = 0 - quotient
	}

	return int(quotient)
}
