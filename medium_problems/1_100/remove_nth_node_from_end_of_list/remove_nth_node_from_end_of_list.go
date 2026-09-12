package remove_nth_node_from_end_of_list

type ListNode struct {
	Val  int
	Next *ListNode
}

func RemoveNthFromEnd(head *ListNode, n int) *ListNode {
	// First pass: length. Counting up front turns "nth from the end" into
	// "node at index count-n from the front", which is a position this list
	// can actually be walked to - a singly linked list has no way back.
	temp := head
	count := 0
	for temp != nil {
		count++
		temp = temp.Next
	}
	if count == 1 {
		// Only node in the list, and n must be 1, so nothing is left.
		return nil
	} else if count == n {
		// Removing the head itself: there is no previous node to relink,
		// so the new head is simply the second node.
		head = head.Next
	} else {
		// Stop one short of the target so temp is the node *before* it.
		// A singly linked list can only unlink a node from its predecessor.
		temp = head
		x := 1
		for x < count-n {
			x++
			temp = temp.Next
		}
		temp.Next = temp.Next.Next
	}
	return head
}
