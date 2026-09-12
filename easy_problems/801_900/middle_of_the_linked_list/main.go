package middleofthelinkedlist

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func middleNode(head *ListNode) *ListNode {
	middle, end := head, head

	// end covers two nodes for every one middle covers, so when end reaches
	// the far end middle has travelled exactly half the distance. The two
	// conditions are both needed: end != nil catches an even-length list,
	// where end steps onto nil, and end.Next != nil catches an odd-length
	// one, where end lands on the last node and end.Next.Next would panic.
	for end != nil && end.Next != nil {
		middle = middle.Next
		end = end.Next.Next
	}
	// For an even count this stops on the second of the two middle nodes,
	// which is what the problem asks for. No length is ever computed.
	return middle
}
