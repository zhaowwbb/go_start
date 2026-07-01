package main

import (
	"fmt"
	"math"
	"testing"
)

// TestCase defines the structure for our division test scenarios
type TestCase struct {
	dividend int
	divisor  int
	expected int
}

func TestDivide(t *testing.T) {
	testCases := []TestCase{
		{dividend: 10, divisor: 3, expected: 3},
		{dividend: 7, divisor: -3, expected: -2},
		{dividend: 0, divisor: 1, expected: 0},
		{dividend: 1, divisor: 1, expected: 1},
		{dividend: -1, divisor: 1, expected: -1},
		// Overflow / Bound Edge Cases
		{dividend: math.MinInt32, divisor: -1, expected: math.MaxInt32}, // Max overflow
		{dividend: math.MinInt32, divisor: 1, expected: math.MinInt32},
		{dividend: math.MaxInt32, divisor: 1, expected: math.MaxInt32},
		{dividend: 100, divisor: 7, expected: 14},
	}

	fmt.Println("Running test cases for Divide Two Integers...")
	fmt.Println("--------------------------------------------------")

	passedCount := 0
	failedCount := 0

	for i, tc := range testCases {
		// Call the implementation exactly once per test case
		// result := Divide(tc.dividend, tc.divisor)
		// result := DivideV2(tc.dividend, tc.divisor)
		// result := DivideV3(tc.dividend, tc.divisor)
		result := DivideV4(tc.dividend, tc.divisor)

		passed := result == tc.expected

		status := "FAILED"
		if passed {
			passedCount++
			status = "PASSED"
		} else {
			failedCount++
			t.Errorf("Test Case %d failed", i+1)
		}

		fmt.Printf("Test Case %d: %s\n", i+1, status)
		fmt.Printf("  Dividend: %d | Divisor: %d\n", tc.dividend, tc.divisor)
		fmt.Printf("  Result:   %d (Expected: %d)\n", result, tc.expected)
		fmt.Println("--------------------------------------------------")
	}

	// Final Summary Report
	fmt.Println("\n================ TEST SUMMARY ================")
	fmt.Printf("Total Test Cases: %d\n", len(testCases))
	fmt.Printf("Passed:           %d\n", passedCount)
	fmt.Printf("Failed:           %d\n", failedCount)
	fmt.Println("==============================================")
}
