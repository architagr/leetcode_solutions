package reverse_linked_list

type ListNode struct {
	Val  int
	Next *ListNode
}

func ReverseList(head *ListNode) *ListNode {

	var tail *ListNode = nil
	for head != nil {
		node := new(ListNode)
		node.Val = head.Val
		node.Next = tail
		tail = node
		head = head.Next
	}
	return tail
}
