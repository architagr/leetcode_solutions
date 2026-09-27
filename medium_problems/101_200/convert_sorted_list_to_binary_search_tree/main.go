package convertsortedlisttobinarysearchtree

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedListToBST(head *ListNode) *TreeNode {
	// A list can't jump to its middle, so copy it into a slice once and reuse
	// the sorted-array build. Costs O(n) space; fast/slow pointers per level
	// would avoid the copy at O(n log n) time.
	nums := listToArray(head)
	return sortedArrayToBST(nums)
}
func sortedArrayToBST(nums []int) *TreeNode {
	n := len(nums)
	if n == 0 {
		return nil
	}
	mid := n / 2
	node := new(TreeNode)
	node.Val = nums[mid]
	node.Left = sortedArrayToBST(nums[:mid])
	node.Right = sortedArrayToBST(nums[mid+1:])
	return node
}

func listToArray(head *ListNode) []int {
	result := make([]int, 0)
	for head != nil {
		result = append(result, head.Val)
		head = head.Next
	}
	return result
}

type ListNode struct {
	Val  int
	Next *ListNode
}
