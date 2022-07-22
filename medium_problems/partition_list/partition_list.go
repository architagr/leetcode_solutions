package partition_list

type ListNode struct {
	Val  int
	Next *ListNode
}

func PartitionList(head *ListNode, x int) *ListNode {
	if x < -100 || x > 100 || head == nil {
		return head
	}
	temp := head
	var newHead *ListNode = nil
	var tempNewHead *ListNode = nil

	var newHead2 *ListNode = nil
	var tempNewHead2 *ListNode = nil

	for temp != nil {
		if temp.Val < x {
			if newHead == nil {
				x := new(ListNode)
				x.Val = temp.Val
				newHead = x
				tempNewHead = newHead
			} else {
				x := new(ListNode)
				x.Val = temp.Val
				tempNewHead.Next = x
				tempNewHead = tempNewHead.Next
			}
		} else {
			if newHead2 == nil {
				x := new(ListNode)
				x.Val = temp.Val
				newHead2 = x
				tempNewHead2 = newHead2
			} else {
				x := new(ListNode)
				x.Val = temp.Val
				tempNewHead2.Next = x
				tempNewHead2 = tempNewHead2.Next
			}
		}
		temp = temp.Next
	}
	if tempNewHead == nil {
		return newHead2
	}
	tempNewHead.Next = newHead2

	return newHead
}
