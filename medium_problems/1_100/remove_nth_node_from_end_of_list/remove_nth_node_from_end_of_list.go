package remove_nth_node_from_end_of_list

type ListNode struct {
	Val  int
	Next *ListNode
}

func RemoveNthFromEnd(head *ListNode, n int) *ListNode {
	temp := head
	count := 0
	for temp != nil {
		count++
		temp = temp.Next
	}
	if count == 1 {
		return nil
	} else if count == n {
		head = head.Next
	} else {
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
