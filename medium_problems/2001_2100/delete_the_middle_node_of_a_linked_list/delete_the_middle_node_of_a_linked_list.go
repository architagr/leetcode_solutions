package delete_the_middle_node_of_a_linked_list

type ListNode struct {
	Val  int
	Next *ListNode
}

func DeleteMiddle(head *ListNode) *ListNode {
	temp := head
	count := 0
	// count number of nodes
	for temp != nil {
		count++
		temp = temp.Next
	}
	// if we just have 1 node then we do not care about n the return will be nil as n has min 1
	if count == 1 {
		return nil
	} else {
		n := count / 2
		temp = head
		x := 1
		// keep moving ahead for count/2 nodes
		for x < n {
			x++
			temp = temp.Next
		}
		// remove node
		temp.Next = temp.Next.Next
	}
	return head
}
