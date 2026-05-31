package main

// ListNode defines the structure for a singly-linked list node.
type ListNode struct {
	Val  int
	Next *ListNode
}

type ListNode2 struct {
	Val  int
	Next *ListNode2
}

type ListNode3 struct {
	Val  int
	Next *ListNode3
}

// addTwoNumbers adds two numbers represented by linked lists and returns the sum as a linked list.
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummyHead := &ListNode{Val: 0}
	current := dummyHead
	carry := 0

	// Continue looping if there are nodes left in l1 OR l2, OR if a carry remains
	for l1 != nil || l2 != nil || carry != 0 {
		val1 := 0
		if l1 != nil {
			val1 = l1.Val
		}

		val2 := 0
		if l2 != nil {
			val2 = l2.Val
		}

		// Calculate total sum for this position
		sum := val1 + val2 + carry

		// Calculate new carry and the digit value to store
		carry = sum / 10
		digit := sum % 10

		// Create a new node and move the tracking pointer forward
		current.Next = &ListNode{Val: digit}
		current = current.Next

		// Move list pointers forward if they haven't reached the end
		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}
	}

	return dummyHead.Next
}

func addTwoNumbersV2(l1 *ListNode, l2 *ListNode) *ListNode {
	dummyHead := &ListNode{Val: 0}
	current := dummyHead
	carry := 0

	for l1 != nil || l2 != nil || carry != 0 {
		val1 := 0
		if l1 != nil {
			val1 = l1.Val
		}
		val2 := 0
		if l2 != nil {
			val2 = l2.Val
		}

		sum := val1 + val2 + carry
		digit := sum % 10
		carry = sum / 10

		current.Next = &ListNode{Val: digit}
		current = current.Next

		if l1 != nil {
			l1 = l1.Next
		}

		if l2 != nil {
			l2 = l2.Next
		}
	}
	return dummyHead.Next
}

func addTwoNumbersV3(l1 *ListNode, l2 *ListNode) *ListNode {
	dummyHead := &ListNode{Val: 0}
	current := dummyHead
	carry := 0

	for l1 != nil || l2 != nil || carry != 0 {
		val1 := 0
		if l1 != nil {
			val1 = l1.Val
		}
		val2 := 0
		if l2 != nil {
			val2 = l2.Val
		}
		sum := val1 + val2 + carry
		digit := sum % 10
		carry = sum / 10

		current.Next = &ListNode{Val: digit}
		current = current.Next

		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}
	}
	return dummyHead.Next
}
