package linked_list_cycle

type ListNode struct {
	Val  int
	Next *ListNode
}

func HasCycle(head *ListNode) bool {
	slow, fast := head, head

	// Same two-pointer walk as 876. The difference is what the gap means:
	// inside a loop fast closes on slow by exactly one node per turn, so if
	// a cycle exists they are guaranteed to land on the same node rather
	// than stepping over each other.
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		// Compare the nodes themselves, not their values. Two nodes can
		// hold the same number without being the same node.
		if slow == fast {
			return true
		}
	}
	// fast reached nil, so the list has an end and nothing loops back.
	return false
}
