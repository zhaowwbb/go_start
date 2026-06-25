package main

import (
	"fmt"
	"testing"
)

func TestRomanToInt(t *testing.T) {
	// Test cases: map[input_roman_string]expected_integer
	testCases := map[string]int{
		// LeetCode Examples
		"III":        3,
		"LVIII":      58,
		"MCMXCIV":    1994,
		"MMMDCCXLIX": 3749,

		// Single Base Symbols
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,

		// Subtractive Forms
		"IV": 4,
		"IX": 9,
		"XL": 40,
		"XC": 90,
		"CD": 400,
		"CM": 900,

		// Boundary Constraints
		"MMMCMXCIX": 3999,
	}

	passedCount := 0
	failedCount := 0

	fmt.Println() // Spacer for clean layout under go test output
	fmt.Printf("%-15s | %-10s | %-10s | %s\n", "INPUT", "EXPECTED", "ACTUAL", "STATUS")
	fmt.Println("------------------------------------------------------------")

	for roman, expected := range testCases {
		// The implementation is invoked exactly once per test case execution
		// actual := RomanToInt(roman)
		// actual := RomanToIntV2(roman)
		actual := RomanToIntV3(roman)

		status := "PASS"
		if actual == expected {
			passedCount++
		} else {
			status = fmt.Sprintf("FAIL (Expected %d, got %d)", expected, actual)
			failedCount++
			t.Errorf("RomanToInt(%s) failed: expected %d, got %d", roman, expected, actual)
		}

		fmt.Printf("%-15s | %-10d | %-10d | %s\n", roman, expected, actual, status)
	}

	fmt.Println("------------------------------------------------------------")
	fmt.Println("Execution Summary:")
	fmt.Printf("Total Test Cases: %d\n", len(testCases))
	fmt.Printf("PASSED          : %d\n", passedCount)
	fmt.Printf("FAILED          : %d\n", failedCount)
	fmt.Println("------------------------------------------------------------")
}
