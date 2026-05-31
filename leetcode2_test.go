package main

import (
	"fmt"
	"reflect"
	"testing"
)

// AddTwoNumbersTestCase defines the structure for our test group
type AddTwoNumbersTestCase struct {
	id       int
	l1       []int
	l2       []int
	expected []int
}

// sliceToLinkedList converts a Go slice into a ListNode linked list.
func sliceToLinkedList(nums []int) *ListNode {
	dummy := &ListNode{}
	curr := dummy
	for _, val := range nums {
		curr.Next = &ListNode{Val: val}
		curr = curr.Next
	}
	return dummy.Next
}

// linkedListToSlice converts a ListNode linked list back into a Go slice.
func linkedListToSlice(node *ListNode) []int {
	var result []int
	curr := node
	for curr != nil {
		result = append(result, curr.Val)
		curr = curr.Next
	}
	return result
}

func TestAddTwoNumbers(t *testing.T) {
	// Test case group containing all scenario data
	testCases := []AddTwoNumbersTestCase{
		{
			id:       1,
			l1:       []int{2, 4, 3},
			l2:       []int{5, 6, 4},
			expected: []int{7, 0, 8},
		},
		{
			id:       2,
			l1:       []int{0},
			l2:       []int{0},
			expected: []int{0},
		},
		{
			id:       3,
			l1:       []int{9, 9, 9, 9, 9, 9, 9},
			l2:       []int{9, 9, 9, 9},
			expected: []int{8, 9, 9, 9, 0, 0, 0, 1},
		},
	}

	// Loop through the test case group
	for _, tc := range testCases {
		// Convert input slices into linked lists
		head1 := sliceToLinkedList(tc.l1)
		head2 := sliceToLinkedList(tc.l2)

		// Call implementation function exactly once
		// actualNodeResult := addTwoNumbers(head1, head2)
		// actualNodeResult := addTwoNumbersV2(head1, head2)
		actualNodeResult := addTwoNumbersV3(head1, head2)

		// Convert result back to slice for printing and verification
		actualSliceResult := linkedListToSlice(actualNodeResult)

		fmt.Printf("Test Case %d:\n", tc.id)
		fmt.Printf("  l1:       %v\n", tc.l1)
		fmt.Printf("  l2:       %v\n", tc.l2)
		fmt.Printf("  Expected: %v\n", tc.expected)
		fmt.Printf("  Actual:   %v\n", actualSliceResult)
		fmt.Println("--------------------------------------------")

		// Verify correctness
		if !reflect.DeepEqual(actualSliceResult, tc.expected) {
			t.Errorf("Test Case %d failed. Expected %v, got %v", tc.id, tc.expected, actualSliceResult)
		}
	}
}
