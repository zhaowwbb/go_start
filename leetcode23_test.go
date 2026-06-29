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

func TestMergeKLists(t *testing.T) {
	// Define the test cases
	testCases := []struct {
		input    [][]int
		expected []int
	}{
		{
			input:    [][]int{{1, 4, 5}, {1, 3, 4}, {2, 6}},
			expected: []int{1, 1, 2, 3, 4, 4, 5, 6},
		},
		{
			input:    [][]int{},
			expected: []int{},
		},
		{
			input:    [][]int{{}},
			expected: []int{},
		},
		{
			input:    [][]int{{1, 3, 5, 7}, {2, 4, 6, 8}},
			expected: []int{1, 2, 3, 4, 5, 6, 7, 8},
		},
		{
			input:    [][]int{{-1, 5, 11}, {}, {6, 10}},
			expected: []int{-1, 5, 6, 10, 11},
		},
	}

	successCount := 0
	failedCount := 0

	t.Log("Running Tests for Merge k Sorted Lists...")
	t.Log("--------------------------------------------------")

	for i, tc := range testCases {
		// Prepare the input: converting slices of slices to linked lists
		var lists []*ListNode
		for _, arr := range tc.input {
			lists = append(lists, createLinkedList(arr))
		}

		// Run implementation
		// mergedHead := mergeKLists(lists)
		// mergedHead := mergeKListsV2(lists)
		mergedHead := mergeKListsV3(lists)

		actual := linkedListToSlice(mergedHead)

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
