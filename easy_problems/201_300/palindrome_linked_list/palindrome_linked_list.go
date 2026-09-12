package palindrome_linked_list

type ListNode struct {
	Val  int
	Next *ListNode
}

func IsPalindrome(head *ListNode) bool {
	// Empty and single-node lists read the same both ways.
	if head == nil || head.Next == nil {
		return true
	}

	// Fast/slow split: when fast falls off the end, slow is on the middle.
	// For an even length slow lands on the first node of the second half,
	// which is what we want - the two halves come out the same length.
	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	// Reverse everything after slow, so the back half can be walked forward.
	second := reverse(slow.Next)

	// Compare inward. The second half is the shorter one when the length is
	// odd, so it is the one that decides when to stop - the odd middle node
	// has no partner and never needs comparing.
	left, right := head, second
	ok := true
	for right != nil {
		if left.Val != right.Val {
			ok = false
			break
		}
		left = left.Next
		right = right.Next
	}

	// Put the list back the way it was found. The caller handed us their
	// list, not a copy, and silently leaving half of it reversed is a
	// surprise they did not ask for.
	slow.Next = reverse(second)
	return ok
}

// reverse relinks a list in place and returns its new head - the same walk
// as 206, which this problem is one half of.
func reverse(head *ListNode) *ListNode {
	var prev *ListNode
	for head != nil {
		next := head.Next
		head.Next = prev
		prev = head
		head = next
	}
	return prev
}
