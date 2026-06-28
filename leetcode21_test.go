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

func TestMergeTwoLists(t *testing.T) {
	type testCase struct {
		list1    []int
		list2    []int
		expected []int
	}

	tests := []testCase{
		{list1: []int{1, 2, 4}, list2: []int{1, 3, 4}, expected: []int{1, 1, 2, 3, 4, 4}},
		{list1: []int{}, list2: []int{}, expected: []int{}},
		{list1: []int{}, list2: []int{0}, expected: []int{0}},
		{list1: []int{1, 5, 9}, list2: []int{2, 3, 4, 6, 7}, expected: []int{1, 2, 3, 4, 5, 6, 7, 9}},
	}

	successCount := 0
	failedCount := 0

	fmt.Println("\nExecuting Test Cases...")
	fmt.Println("----------------------------------------")

	for i, tc := range tests {
		// Create fresh copies of lists for each single-run iteration
		l1 := sliceToLinkedList(tc.list1)
		l2 := sliceToLinkedList(tc.list2)

		// Call implementation exactly once per case
		// resultHead := MergeTwoLists(l1, l2)
		// resultHead := MergeTwoListsV2(l1, l2)
		resultHead := MergeTwoListsV3(l1, l2)

		resultSlice := linkedListToSlice(resultHead)

		if (len(tc.expected) == 0 && len(resultSlice) == 0) || reflect.DeepEqual(resultSlice, tc.expected) {
			fmt.Printf("Test Case %d: PASSED\n", i+1)
			successCount++
		} else {
			fmt.Printf("Test Case %d: FAILED (L1: %v, L2: %v | Expected: %v | Got: %v)\n",
				i+1, tc.list1, tc.list2, tc.expected, resultSlice)
			failedCount++
			t.Errorf("Test Case %d failed", i+1)
		}
	}

	fmt.Println("----------------------------------------")
	fmt.Printf("SUMMARY: %d passed, %d failed.\n\n", successCount, failedCount)
}
