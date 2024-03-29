package merge_k_sorted_lists

import "math"

type ListNode struct {
	Val  int
	Next *ListNode
}

func MergeKLists(lists []*ListNode) *ListNode {
	var head *ListNode = nil
	var temp *ListNode = nil

	n, x := len(lists), 0

	for x < n {
		x = 0
		min := math.MaxInt32
		minIndex := n + 1
		// find min from current node of the entire list
		for i := 0; i < n; i++ {
			if lists[i] != nil {
				if lists[i].Val < min {
					min = lists[i].Val
					minIndex = i
				}
			} else {
				x++
			}
		}
		// if conlete list was not nil then push the min to the new linked list
		if x < n {
			if head == nil {
				head = new(ListNode)
				head.Val = min
				temp = head
			} else {
				temp.Next = new(ListNode)
				temp = temp.Next
				temp.Val = min
			}
			// move the node of the minIndex to next
			lists[minIndex] = lists[minIndex].Next
		}
	}

	return head
}
