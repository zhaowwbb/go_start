package main

import (
	"fmt"
	"reflect"
	"testing"
)

// helper to convert a slice into a Linked List
func sliceToLinkedList(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}
	dummy := &ListNode{}
	curr := dummy
	for _, val := range nums {
		curr.Next = &ListNode{Val: val}
		curr = curr.Next
	}
	return dummy.Next
}

// helper to convert a Linked List back into a slice
func linkedListToSlice(head *ListNode) []int {
	var result []int
	curr := head
	for curr != nil {
		result = append(result, curr.Val)
		curr = curr.Next
	}
	return result
}

func TestRemoveNthFromEnd(t *testing.T) {
	// Define test cases
	type testCase struct {
		input    []int
		n        int
		expected []int
	}

	tests := []testCase{
		{input: []int{1, 2, 3, 4, 5}, n: 2, expected: []int{1, 2, 3, 5}}, // Standard case
		{input: []int{1}, n: 1, expected: []int{}},                       // Single node removal
		{input: []int{1, 2}, n: 1, expected: []int{1}},                   // Remove last node
		{input: []int{1, 2}, n: 2, expected: []int{2}},                   // Remove first node (head)
		{input: []int{1, 2, 3}, n: 3, expected: []int{2, 3}},             // Remove head from longer list
	}

	successCount := 0
	failedCount := 0

	fmt.Println("\nExecuting Test Cases...")
	fmt.Println("----------------------------------------")

	for i, tc := range tests {
		// Rebuild fresh linked list per iteration
		head := sliceToLinkedList(tc.input)

		// Call the implementation exactly once per test case
		// resultHead := RemoveNthFromEnd(head, tc.n)
		// resultHead := RemoveNthFromEndV2(head, tc.n)
		resultHead := RemoveNthFromEndV3(head, tc.n)

		// Convert result back to standard slice for matching
		resultSlice := linkedListToSlice(resultHead)

		// Check results using reflect.DeepEqual (handles empty slices/nil well)
		if (len(tc.expected) == 0 && len(resultSlice) == 0) || reflect.DeepEqual(resultSlice, tc.expected) {
			fmt.Printf("Test Case %d: PASSED\n", i+1)
			successCount++
		} else {
			fmt.Printf("Test Case %d: FAILED (Input: %v, n: %d | Expected: %v | Got: %v)\n",
				i+1, tc.input, tc.n, tc.expected, resultSlice)
			failedCount++
			t.Errorf("Test Case %d failed", i+1) // Signals native 'go test' runner of failure
		}
	}

	fmt.Println("----------------------------------------")
	fmt.Printf("SUMMARY: %d passed, %d failed.\n\n", successCount, failedCount)
}
