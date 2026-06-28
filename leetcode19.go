package main

// ListNode defines a node for a singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// RemoveNthFromEnd removes the nth node from the end of the list and returns its head.
// It uses a dummy node and two pointers (fast and slow) to achieve O(N) time and O(1) space.
func RemoveNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{Val: 0, Next: head}
	fast := dummy
	slow := dummy

	// Advance fast pointer so that the gap between fast and slow is n nodes
	for i := 0; i <= n; i++ {
		if fast == nil {
			return head // Edge case guard
		}
		fast = fast.Next
	}

	// Move fast to the end, maintaining the gap
	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}

	// slow.Next is the node to be deleted
	if slow.Next != nil {
		slow.Next = slow.Next.Next
	}

	return dummy.Next
}

func RemoveNthFromEndV2(head *ListNode, n int) *ListNode {
	dummy := &ListNode{Val: 0, Next: head}
	slow := dummy
	fast := dummy
	for i := 0; i <= n; i++ {
		if fast == nil {
			return head
		}
		fast = fast.Next
	}

	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}
	slow.Next = slow.Next.Next

	return dummy.Next
}

func RemoveNthFromEndV3(head *ListNode, n int) *ListNode {
	dummy := &ListNode{Val: 0, Next: head}
	fast := dummy
	slow := dummy
	for i := 0; i <= n; i++ {
		if fast == nil {
			return head
		}
		fast = fast.Next
	}
	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}
	slow.Next = slow.Next.Next

	return dummy.Next
}
