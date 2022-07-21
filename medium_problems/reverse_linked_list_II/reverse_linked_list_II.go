package reverse_linked_list_II

type ListNode struct {
	Val  int
	Next *ListNode
}

func ReverseBetween(head *ListNode, left int, right int) *ListNode {
	if left == right {
		return head
	}
	tempHead := head
	var leftNode *ListNode = nil
	var rightNode *ListNode = nil
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

	rightNode = tempHead.Next
	tempHead.Next = nil
	leftSide = reverse(leftSide)
	if leftNode == nil {
		head = leftSide
	} else {
		leftNode.Next = leftSide
	}

	for leftSide.Next != nil {
		leftSide = leftSide.Next
	}
	leftSide.Next = rightNode
	return head
}

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
