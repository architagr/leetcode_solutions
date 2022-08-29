package delete_node_in_a_linked_list

type ListNode struct {
	Val  int
	Next *ListNode
}

func DeleteNode(node *ListNode) {

	dummy := node

	for dummy.Next != nil {
		dummy.Val = dummy.Next.Val
		if dummy.Next.Next == nil {
			dummy.Next = nil
		} else {
			dummy = dummy.Next
		}

	}

}
