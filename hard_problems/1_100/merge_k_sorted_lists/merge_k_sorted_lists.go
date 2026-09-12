package merge_k_sorted_lists

import "container/heap"

type ListNode struct {
	Val  int
	Next *ListNode
}

// nodeHeap orders list heads by value, so the smallest unplaced node across
// all k lists is always at index 0.
//
// Only one node per list is ever in the heap. A list's second node becomes a
// candidate exactly when its first is placed, so the heap stays at size k
// however many nodes the lists hold between them.
type nodeHeap []*ListNode

func (h nodeHeap) Len() int           { return len(h) }
func (h nodeHeap) Less(i, j int) bool { return h[i].Val < h[j].Val }
func (h nodeHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *nodeHeap) Push(x any)        { *h = append(*h, x.(*ListNode)) }
func (h *nodeHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func MergeKLists(lists []*ListNode) *ListNode {
	h := &nodeHeap{}
	for _, l := range lists {
		// Empty lists contribute nothing and must not be pushed - Less would
		// dereference a nil node.
		if l != nil {
			*h = append(*h, l)
		}
	}
	// Heapify once rather than pushing k times: O(k) instead of O(k log k).
	heap.Init(h)

	// A dummy head, as in 21, so the loop never has to ask which node is first.
	dummy := &ListNode{}
	tail := dummy

	for h.Len() > 0 {
		node := heap.Pop(h).(*ListNode)
		// Splice the existing node in rather than copying its value into a
		// new one. Nothing is allocated per element.
		tail.Next = node
		tail = node
		if node.Next != nil {
			heap.Push(h, node.Next)
		}
	}

	// The last node popped is the largest, so it is the tail of its own list
	// and its Next is already nil - anything after it would have been pushed
	// and popped later. This is belt and braces: splicing reuses nodes whose
	// Next pointers still refer to their original lists, and a merged list
	// that is not explicitly terminated is the way that becomes a cycle.
	tail.Next = nil
	return dummy.Next
}
