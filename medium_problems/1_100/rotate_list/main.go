package rotatelist

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	n, lastNode := length(head)
	if n < 2 || k == 0 {
		return head
	}
	k = k % n
	temp := head
	for i := 1; i < n-k; i++ {
		temp = temp.Next
	}
	lastNode.Next = head
	head = temp.Next
	temp.Next = nil
	return head
}

func length(head *ListNode) (count int, lastNode *ListNode) {
	temp := head
	count = 0
	for temp != nil {
		lastNode = temp
		temp = temp.Next
		count++
	}
	return
}
