package reverse_linked_list_II

type ListNode struct {
	Val  int
	Next *ListNode
}

func ReverseBetween(head *ListNode, left int, right int) *ListNode {
	// A one-node run is already reversed.
	if left == right {
		return head
	}
	tempHead := head
	var leftNode *ListNode = nil
	var rightNode *ListNode = nil
	// Walk to position left, keeping leftNode one step behind. That
	// trailing pointer is what the reversed run gets stitched back onto,
	// and it stays nil when left == 1 because there is nothing before the
	// head to stitch to.
	count := 0
	for tempHead != nil {
		count++
		if count == left {
			break
		}
		leftNode = tempHead
		tempHead = tempHead.Next
	}
	leftSide := tempHead
	for tempHead != nil {
		if count == right {
			break
		}
		count++
		tempHead = tempHead.Next
	}

	// Save what follows the run, then cut the run loose. Detaching it means
	// the plain 206 reversal can be reused as-is, with no notion of where
	// to stop built into it.
	rightNode = tempHead.Next
	tempHead.Next = nil
	leftSide = reverse(leftSide)
	// Reattach the front. When left == 1 the run started at the head, so
	// the reversed run's head becomes the list's head.
	if leftNode == nil {
		head = leftSide
	} else {
		leftNode.Next = leftSide
	}

	// Reattach the back. After reversing, what was the run's first node is
	// now its last, so walk to the end of the run to find it.
	for leftSide.Next != nil {
		leftSide = leftSide.Next
	}
	leftSide.Next = rightNode
	return head
}

// reverse is the 206 walk, unchanged: point each node at its predecessor
// and return the node that ends up first.
func reverse(head *ListNode) *ListNode {
	tempHead := head
	var prev *ListNode = nil
	var next *ListNode = nil

	for tempHead != nil {
		next = tempHead.Next
		tempHead.Next = prev
		prev = tempHead
		tempHead = next
	}
	return prev
}
