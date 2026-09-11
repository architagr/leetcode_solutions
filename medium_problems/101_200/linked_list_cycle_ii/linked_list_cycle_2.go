package linked_list_cycle_2

type ListNode struct {
	Val  int
	Next *ListNode
}

func DetectCycle(head *ListNode) *ListNode {
	mapN := make(map[*ListNode]bool)
	node := head
	for node != nil {
		if _, ok := mapN[node]; ok {
			return node
		}
		mapN[node] = true
		node = node.Next
	}
	return nil
}
