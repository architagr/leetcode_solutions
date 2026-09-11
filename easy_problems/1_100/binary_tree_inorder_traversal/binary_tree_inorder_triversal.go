package binary_tree_inorder_triversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func InorderTraversal(root *TreeNode) []int {
	arr := make([]int, 0, 100)
	return traversal(root, arr)
}
func traversal(head *TreeNode, arr []int) []int {
	if head == nil {
		return arr
	}

	arr = traversal(head.Left, arr)
	arr = append(arr, head.Val)
	arr = traversal(head.Right, arr)

	return arr
}
