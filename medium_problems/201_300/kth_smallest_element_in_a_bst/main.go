package kthsmallestelementinabst

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func kthSmallest(root *TreeNode, k int) int {
	// In-order visits a BST in sorted order, so the kth value appended is the
	// answer. The walk could stop there; this version collects everything.
	arr := make([]int, 0, 10_000)
	arr = inOrder(root, arr)
	return arr[k-1]
}
func inOrder(node *TreeNode, arr []int) []int {
	if node == nil {
		return arr
	}

	arr = inOrder(node.Left, arr)
	arr = append(arr, node.Val)
	arr = inOrder(node.Right, arr)
	return arr
}
