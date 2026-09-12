package reverse_linked_list

type ListNode struct {
	Val  int
	Next *ListNode
}

func ReverseList(head *ListNode) *ListNode {
	// nil is the correct final Next for the current head: it ends up last in
	// the reversed list, and the last node points at nil. Starting here is
	// what removes the need for a special case on the old head.
	var prev *ListNode = nil
	for head != nil {
		// Save the rest of the list before the next line destroys the only
		// pointer to it. Without this the walk cannot continue.
		next := head.Next
		head.Next = prev
		// Step both pointers forward, prev first - it has to become the node
		// just finished before head moves off it.
		prev = head
		head = next
	}
	// prev, not head: head is nil here. prev is the original tail, which is
	// the head of the reversed list.
	return prev
}
