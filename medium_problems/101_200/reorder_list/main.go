package reorderlist

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// reorderList weaves the list's two halves together. The whole function is
// three techniques the series has already covered, run in sequence: find the
// middle, reverse the back half, then merge the two alternately. Nothing here
// is new - what is new is noticing that the target order IS that merge.
func reorderList(head *ListNode) {
	mid := getMid(head)
	head2 := mid.Next

	// Cut, so the first half terminates instead of running into the second.
	// Without this the merge walks into nodes it has already placed.
	mid.Next = nil
	head2 = reverseList(head2)

	// mergeList relinks in place, so the caller's head still points at the
	// right node - which is why this function can return nothing.
	head = mergeList(head, head2)
}

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	// The 206 walk, with next hoisted out of the loop rather than declared
	// inside it.
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
	// t walks the first half, a2 the reversed second half, and t1/t2 hold
	// each one's successor across the relinking - the same "save it before
	// you overwrite it" problem as 206, now with two lists to keep alive.
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
	// 876's fast/slow split. Testing fast.Next and fast.Next.Next rather
	// than fast stops slow on the LAST node of the first half, which is the
	// node that has to be cut - not the first node of the second half.
	slow, fast := head, head

	for fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	return slow
}
