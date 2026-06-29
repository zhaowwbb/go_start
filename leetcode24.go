package main

// ListNode defines a node for a singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// swapPairs swaps every two adjacent nodes in a linked list and returns its head.
func swapPairs(head *ListNode) *ListNode {
	// Create a dummy node to handle edge cases easily at the head of the list
	dummy := &ListNode{Val: 0, Next: head}
	prev := dummy

	// Ensure there are at least two nodes left to swap
	for prev.Next != nil && prev.Next.Next != nil {
		first := prev.Next
		second := prev.Next.Next

		// Perform the swap by altering pointers
		first.Next = second.Next
		second.Next = first
		prev.Next = second

		// Advance the prev pointer forward by two nodes for the next iteration
		prev = first
	}

	return dummy.Next
}

func swapPairsV2(head *ListNode) *ListNode {
	dummy := &ListNode{Val: 0, Next: head}
	prev := dummy

	for prev.Next != nil && prev.Next.Next != nil {
		first := prev.Next
		second := prev.Next.Next

		first.Next = second.Next
		prev.Next = second
		second.Next = first

		prev = first
	}
	return dummy.Next
}

func swapPairsV3(head *ListNode) *ListNode {
	dummy := &ListNode{Val: 0, Next: head}
	prev := dummy

	for prev.Next != nil && prev.Next.Next != nil {
		first := prev.Next
		second := prev.Next.Next

		first.Next = second.Next
		prev.Next = second
		second.Next = first

		prev = first
	}
	return dummy.Next
}
