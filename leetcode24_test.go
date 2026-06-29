package main

import (
	"testing"
)

// Helper to convert a standard slice into a linked list
func createLinkedList(nums []int) *ListNode {
	dummy := &ListNode{}
	curr := dummy
	for _, val := range nums {
		curr.Next = &ListNode{Val: val}
		curr = curr.Next
	}
	return dummy.Next
}

// Helper to convert a linked list back into a standard slice
func linkedListToSlice(head *ListNode) []int {
	var result []int
	for head != nil {
		result = append(result, head.Val)
		head = head.Next
	}
	return result
}

// Helper to compare two integer slices
func slicesEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func TestSwapPairs(t *testing.T) {
	// Define test cases
	testCases := []struct {
		input    []int
		expected []int
	}{
		{
			input:    []int{1, 2, 3, 4},
			expected: []int{2, 1, 4, 3},
		},
		{
			input:    []int{},
			expected: []int{},
		},
		{
			input:    []int{1},
			expected: []int{1},
		},
		{
			input:    []int{1, 2, 3},
			expected: []int{2, 1, 3},
		},
		{
			input:    []int{7, 9, 2, 8, 3},
			expected: []int{9, 7, 8, 2, 3},
		},
	}

	successCount := 0
	failedCount := 0

	t.Log("Running Tests for Swap Nodes in Pairs...")
	t.Log("--------------------------------------------------")

	for i, tc := range testCases {
		// Convert input array into a working linked list
		head := createLinkedList(tc.input)

		// Run implementation
		// swappedHead := swapPairs(head)
		// swappedHead := swapPairsV2(head)
		swappedHead := swapPairsV3(head)

		actual := linkedListToSlice(swappedHead)

		// Verification
		if slicesEqual(actual, tc.expected) {
			t.Logf("Test Case %d: PASSED", i+1)
			successCount++
		} else {
			t.Errorf("Test Case %d: FAILED\n  Expected: %v\n  Got:      %v", i+1, tc.expected, actual)
			failedCount++
		}
	}

	t.Log("--------------------------------------------------")
	t.Logf("Total Successes: %d", successCount)
	t.Logf("Total Failures:  %d", failedCount)
}
