package main

import (
	"reflect"
	"testing"
)

type TestCase struct {
	input    []int
	k        int
	expected []int
}

func TestReverseKGroup(t *testing.T) {
	testCases := []TestCase{
		{input: []int{1, 2, 3, 4, 5}, k: 2, expected: []int{2, 1, 4, 3, 5}},
		{input: []int{1, 2, 3, 4, 5}, k: 3, expected: []int{3, 2, 1, 4, 5}},
		{input: []int{1, 2, 3, 4, 5}, k: 1, expected: []int{1, 2, 3, 4, 5}},
		{input: []int{1}, k: 1, expected: []int{1}},
		{input: []int{}, k: 3, expected: []int{}},
		{input: []int{1, 2}, k: 3, expected: []int{1, 2}},             // Length less than k
		{input: []int{1, 2, 3, 4}, k: 2, expected: []int{2, 1, 4, 3}}, // Exact multiple of k
	}

	t.Log("Running LeetCode 25: Reverse Nodes in k-Group Tests...\n")

	passedCount := 0
	failedCount := 0

	for i, tc := range testCases {
		head := buildLinkedList(tc.input)
		// resultHead := reverseKGroup(head, tc.k)
		// resultHead := reverseKGroupV2(head, tc.k)
		// resultHead := reverseKGroupV3(head, tc.k)
		resultHead := reverseKGroupV4(head, tc.k)

		resultSlice := linkedListToSlice(resultHead)

		if reflect.DeepEqual(resultSlice, tc.expected) {
			t.Logf("✅ Test %d Passed! (k=%d) -> Input: %v | Output: %v", i+1, tc.k, tc.input, resultSlice)
			passedCount++
		} else {
			t.Errorf("❌ Test %d Failed! (k=%d)\n   Input:    %v\n   Expected: %v\n   Got:      %v", i+1, tc.k, tc.input, tc.expected, resultSlice)
			failedCount++
		}
	}

	// Print Summary Dashboard
	t.Logf("\n========================================")
	t.Logf("TEST RUN SUMMARY")
	t.Logf("========================================")
	t.Logf("Total Test Cases: %d", len(testCases))
	t.Logf("Passed:           %d / %d", passedCount, len(testCases))
	t.Logf("Failed:           %d / %d", failedCount, len(testCases))
	t.Logf("========================================")
}
