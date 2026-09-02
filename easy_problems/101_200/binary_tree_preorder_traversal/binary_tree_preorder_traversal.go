package binary_tree_preorder_traversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func PreorderTraversal(root *TreeNode) []int {
	arr := make([]int, 0)
	return traversal(root, arr)
}

func traversal(A *TreeNode, arr []int) []int {
	if A == nil {
		return arr
	}
	arr = append(arr, A.Val)
	arr = traversal(A.Left, arr)

	arr = traversal(A.Right, arr)
	return arr
}
