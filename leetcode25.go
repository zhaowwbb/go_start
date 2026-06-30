package main

import "fmt"

// ListNode defines a node for the singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// reverseKGroup reverses the nodes of a linked list k at a time and returns its head.
func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil || k == 1 {
		return head
	}

	// Dummy node to handle edge cases at the head of the list easily
	dummy := &ListNode{Val: 0, Next: head}
	groupPrev := dummy

	for {
		// Find the k-th node from groupPrev
		kth := getKthNode(groupPrev, k)
		if kth == nil {
			break // Remaining nodes are fewer than k, leave them as they are
		}

		groupNext := kth.Next

		// Reverse the current k-group
		prev := kth.Next
		curr := groupPrev.Next
		for curr != groupNext {
			temp := curr.Next
			curr.Next = prev
			prev = curr
			curr = temp
		}

		// Connect the reversed group back into the main list
		temp := groupPrev.Next
		groupPrev.Next = kth
		groupPrev = temp
	}

	return dummy.Next
}

// getKthNode navigates k steps forward and returns the node found.
func getKthNode(curr *ListNode, k int) *ListNode {
	for curr != nil && k > 0 {
		curr = curr.Next
		k--
	}
	return curr
}

// buildLinkedList is a helper to convert a slice into a linked list.
func buildLinkedList(nums []int) *ListNode {
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

// linkedListToSlice is a helper to convert a linked list back into a slice for comparison.
func linkedListToSlice(head *ListNode) []int {
	result := []int{}
	curr := head
	for curr != nil {
		result = append(result, curr.Val)
		curr = curr.Next
	}
	return result
}

func main() {
	fmt.Println("To run tests, please execute: go test")
}

func reverseKGroupV2(head *ListNode, k int) *ListNode {
	dummy := &ListNode{Val: 0, Next: head}
	prev := dummy
	first := head
	current := first
	var nodes []*ListNode
	for current != nil {
		nodes = append(nodes, current)
		current = current.Next
	}

	n := len(nodes)

	fmt.Println("List size:", n, ", k=", k)
	if n < k {
		return head
	}

	i := 0

	for i < n {
		for j := k; j > 0; j-- {
			pos := i + j - 1
			fmt.Println("i=", i, "j=", j, "pos=", pos)
			prev.Next = nodes[pos]
			prev = prev.Next
		}
		i += k
		if i+k > n {
			break
		}
	}

	fmt.Println("i=", i)

	if i < n {
		prev.Next = nodes[i]
	} else {
		prev.Next = nil
	}
	fmt.Println("reverse k list")

	return dummy.Next
}

func reverseKGroupV3(head *ListNode, k int) *ListNode {
	dummy := &ListNode{Val: 0, Next: head}
	groupPre := dummy

	for {

		//get kth node
		kthNode := groupPre
		i := 0
		for kthNode != nil && i < k {
			kthNode = kthNode.Next
			i++
		}
		if kthNode == nil {
			break
		}

		// fmt.Println("kthNode, val:", kthNode.Val, ",i=", i)

		groupNext := kthNode.Next
		pre := kthNode.Next
		current := groupPre.Next

		//reverse k node
		for current != groupNext {
			temp := current.Next
			current.Next = pre
			pre = current
			current = temp
		}

		temp := groupPre.Next
		groupPre.Next = kthNode
		groupPre = temp
	}

	return dummy.Next
}

func reverseKGroupV4(head *ListNode, k int) *ListNode {
	dummy := &ListNode{Val: 0, Next: head}
	groupPre := dummy

	for {
		//find k th node
		kthNode := groupPre
		i := 0
		for kthNode != nil && i < k {
			kthNode = kthNode.Next
			i++
		}
		//less than k
		if kthNode == nil {
			break
		}
		//reverse kth node
		groupNext := kthNode.Next
		pre := kthNode.Next
		current := groupPre.Next
		for current != groupNext {
			temp := current.Next
			current.Next = pre
			pre = current
			current = temp
		}

		temp := groupPre.Next
		groupPre.Next = kthNode
		groupPre = temp
		fmt.Println("groupPre, val:", groupPre.Val)

	}
	return dummy.Next
}
