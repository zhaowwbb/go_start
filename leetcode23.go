package main

import (
	"container/heap"
)

// ListNode defines a node for a singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// MinHeap implements heap.Interface and holds *ListNode items.
type MinHeap []*ListNode

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i].Val < h[j].Val }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(*ListNode))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// mergeKLists merges k sorted linked lists into one sorted linked list.
func mergeKLists(lists []*ListNode) *ListNode {
	h := &MinHeap{}
	heap.Init(h)

	// Step 1: Push the head of each non-empty list into the heap
	for _, list := range lists {
		if list != nil {
			heap.Push(h, list)
		}
	}

	dummy := &ListNode{}
	current := dummy

	// Step 2: Extract the minimum node and push its next node back into the heap
	for h.Len() > 0 {
		node := heap.Pop(h).(*ListNode)
		current.Next = node
		current = current.Next

		if node.Next != nil {
			heap.Push(h, node.Next)
		}
	}

	return dummy.Next
}

type MinHeapV2 []*ListNode

func (h MinHeapV2) Len() int { return len(h) }

func (h MinHeapV2) Less(a, b int) bool { return h[a].Val < h[b].Val }

func (h MinHeapV2) Swap(a, b int) { h[a], h[b] = h[b], h[a] }

func (h *MinHeapV2) Push(x interface{}) {
	*h = append(*h, x.(*ListNode))
}

func (h *MinHeapV2) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func mergeKListsV2(lists []*ListNode) *ListNode {
	h := &MinHeapV2{}
	heap.Init(h)

	for _, list := range lists {
		if list != nil {
			heap.Push(h, list)
		}
	}

	dummy := &ListNode{}
	current := dummy

	for h.Len() > 0 {
		node := heap.Pop(h).(*ListNode)
		current.Next = node
		current = current.Next
		if node.Next != nil {
			heap.Push(h, node.Next)
		}
	}

	return dummy.Next
}

type ListNodeV3 struct {
	Val  int
	Next *ListNodeV3
}

type MinHeapV3 []*ListNode

func (h MinHeapV3) Len() int           { return len(h) }
func (h MinHeapV3) Less(i, j int) bool { return h[i].Val < h[j].Val }
func (h MinHeapV3) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *MinHeapV3) Push(x interface{}) {
	*h = append(*h, x.(*ListNode))
}
func (h *MinHeapV3) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func mergeKListsV3(lists []*ListNode) *ListNode {
	h := &MinHeapV3{}
	heap.Init(h)

	dummy := &ListNode{}
	current := dummy

	for _, list := range lists {
		if list != nil {
			heap.Push(h, list)
		}
	}

	for h.Len() > 0 {
		node := heap.Pop(h).(*ListNode)
		current.Next = node
		current = current.Next
		if node.Next != nil {
			heap.Push(h, node.Next)
		}
	}

	return dummy.Next
}
