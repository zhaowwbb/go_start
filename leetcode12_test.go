package main

import (
	"fmt"
	"testing"
)

func TestIntToRoman(t *testing.T) {
	// Test cases: map[input_integer]expected_roman_string
	testCases := map[int]string{
		// LeetCode Examples
		3749: "MMMDCCXLIX",
		58:   "LVIII",
		1994: "MCMXCIV",

		// Boundary Constraints
		1:    "I",
		3999: "MMMCMXCIX",

		// Subtractive Forms
		4:   "IV",
		9:   "IX",
		40:  "XL",
		90:  "XC",
		400: "CD",
		900: "CM",

		// Base Symbols
		5:    "V",
		10:   "X",
		50:   "L",
		100:  "C",
		500:  "D",
		1000: "M",
	}

	passedCount := 0
	failedCount := 0

	fmt.Println() // Spacer for clean layout under go test output
	fmt.Printf("%-10s | %-15s | %-15s | %s\n", "INPUT", "EXPECTED", "ACTUAL", "STATUS")
	fmt.Println("------------------------------------------------------------")

	for num, expected := range testCases {
		// The implementation is invoked exactly once per test case execution
		// actual := IntToRoman(num)
		// actual := IntToRomanV2(num)
		actual := IntToRomanV3(num)

		status := "PASS"
		if actual == expected {
			passedCount++
		} else {
			status = fmt.Sprintf("FAIL (Expected %s, got %s)", expected, actual)
			failedCount++
			t.Errorf("IntToRoman(%d) failed: expected %s, got %s", num, expected, actual)
		}

		fmt.Printf("%-10d | %-15s | %-15s | %s\n", num, expected, actual, status)
	}

	fmt.Println("------------------------------------------------------------")
	fmt.Println("Execution Summary:")
	fmt.Printf("Total Test Cases: %d\n", len(testCases))
	fmt.Printf("PASSED          : %d\n", passedCount)
	fmt.Printf("FAILED          : %d\n", failedCount)
	fmt.Println("------------------------------------------------------------")
}
