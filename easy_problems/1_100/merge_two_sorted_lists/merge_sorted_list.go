package merge_sorted_list

type ListNode struct {
	Val  int
	Next *ListNode
}

func MergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	var mainHead *ListNode = new(ListNode)
	if list1 == nil && list2 == nil {
		return nil
	}
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}
	var currentNode *ListNode = mainHead
	//create start on list1
	if list2.Val <= list1.Val {
		currentNode.Val = list2.Val
		currentNode.Next = nil

		mainHead = currentNode
		list2 = list2.Next

	} else {
		currentNode.Val = list1.Val
		currentNode.Next = nil
		mainHead = currentNode
		list1 = list1.Next
	}

	for list2 != nil && list1 != nil {

		var newNode *ListNode = new(ListNode)
		if list2.Val <= list1.Val {
			newNode.Val = list2.Val
			currentNode.Next = newNode
			list2 = list2.Next

		} else {
			newNode.Val = list1.Val
			currentNode.Next = newNode
			list1 = list1.Next
		}
		currentNode = currentNode.Next
	}
	if list1 == nil && list2 != nil {
		currentNode.Next = list2
	}
	if list2 == nil && list1 != nil {
		currentNode.Next = list1
	}

	return mainHead
}
