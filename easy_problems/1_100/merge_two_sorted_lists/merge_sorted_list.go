package merge_sorted_list

type ListNode struct {
	Val  int
	Next *ListNode
}

func MergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	// A throwaway node to hang the result off. It means the loop never has
	// to ask "is this the first node?" - tail is always something real to
	// append to, so there is no separate case for picking the head.
	dummy := &ListNode{}
	tail := dummy

	for list1 != nil && list2 != nil {
		// <= rather than < keeps equal values in their original order.
		if list1.Val <= list2.Val {
			tail.Next = list1
			list1 = list1.Next
		} else {
			tail.Next = list2
			list2 = list2.Next
		}
		tail = tail.Next
	}

	// One list is empty now. Everything left in the other is already sorted
	// and already larger than everything placed, so it gets attached whole
	// rather than walked node by node.
	if list1 != nil {
		tail.Next = list1
	} else {
		tail.Next = list2
	}

	return dummy.Next
}
