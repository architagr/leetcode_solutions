package reorderlist

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func reorderList(head *ListNode) {
	mid := getMid(head)
	head2 := mid.Next

	mid.Next = nil
	head2 = reverseList(head2)

	head = mergeList(head, head2)
}

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var p, c, n *ListNode = nil, head, head.Next

	for c != nil {
		c.Next = p
		p = c
		c = n
		if n != nil {
			n = n.Next
		}
	}
	return p
}

func mergeList(a1, a2 *ListNode) *ListNode {
	if a2 == nil {
		return a1
	}
	t, t1, t2 := a1, a1.Next, a2.Next

	for a2 != nil {
		t.Next = a2
		a2.Next = t1
		t = t1
		a2 = t2
		if t1 != nil {
			t1 = t1.Next
		}

		if t2 != nil {
			t2 = t2.Next
		}
	}
	return a1
}
func getMid(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	slow, fast := head, head

	for fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	return slow
}
