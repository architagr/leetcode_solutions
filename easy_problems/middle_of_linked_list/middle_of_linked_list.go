package middle_of_linked_list

type ListNode struct {
	Val  int
	Next *ListNode
}

func MiddleNode(head *ListNode) *ListNode {
	var x ListNode = *head
	var middle *ListNode = nil
	count := 0

	for head != nil {
		count++
		head = head.Next
	}
	if count < 2 {
		return &x
	}
	middle = &x
	for i := 0; i < (count / 2); i++ {
		x = *x.Next
	}
	return middle
}
