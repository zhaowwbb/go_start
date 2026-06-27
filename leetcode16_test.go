package main

import (
	"fmt"
	"testing"
)

// Define a struct to match the Python test case dictionary structure
type TestCase struct {
	name     string
	nums     []int
	target   int
	expected int
}

func TestThreeSumClosestSuite(t *testing.T) {
	testCases := []TestCase{
		{name: "Standard LeetCode Example", nums: []int{-1, 2, 1, -4}, target: 1, expected: 2},
		{name: "Exact Match Available", nums: []int{0, 0, 0}, target: 1, expected: 0},
		{name: "All Negative Numbers", nums: []int{-1, -2, -3, -4}, target: -5, expected: -6},
		{name: "Multiple Duplicate Elements", nums: []int{1, 1, 1, 1}, target: 100, expected: 3},
		{name: "Closest Sum is Greater than Target", nums: []int{1, 2, 3, 4}, target: 5, expected: 6},
		{name: "Large Target and Elements", nums: []int{10, 20, 30, 40, 50}, target: 55, expected: 60},
		// Intentionally adding a failing test case to demonstrate failure counting logic
		// {name: "Deliberate Failure Case", nums: []int{1, 1, 1}, target: 3, expected: 999},
	}

	passedCount := 0
	failedCount := 0

	fmt.Println("============================================================")
	fmt.Printf("%35s\n", "RUNNING 3SUM CLOSEST TEST SUITE")
	fmt.Println("============================================================")

	for i, tc := range testCases {
		// The implementation is called exactly ONCE per test case here
		// actual := ThreeSumClosest(tc.nums, tc.target)
		// actual := ThreeSumClosestV2(tc.nums, tc.target)
		actual := ThreeSumClosestV3(tc.nums, tc.target)

		if actual == tc.expected {
			passedCount++
			fmt.Printf("Test %d: %-35s | ✅ PASSED (Result: %d)\n", i+1, tc.name, actual)
		} else {
			failedCount++
			fmt.Printf("Test %d: %-35s | ❌ FAILED (Expected %d, got %d)\n", i+1, tc.name, tc.expected, actual)
			// Inform Go's test runner that the test suite technically failed
			t.Errorf("%s failed", tc.name)
		}
	}

	successRate := (float64(passedCount) / float64(len(testCases))) * 100

	fmt.Println("------------------------------------------------------------")
	fmt.Printf("TOTAL PASSED: %d\n", passedCount)
	fmt.Printf("TOTAL FAILED: %d\n", failedCount)
	fmt.Printf("SUCCESS RATE: %.1f%%\n", successRate)
	fmt.Println("============================================================")
}
