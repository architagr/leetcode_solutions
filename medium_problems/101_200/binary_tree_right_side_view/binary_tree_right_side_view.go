package binary_tree_right_side_view

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func BinaryTreeRightSideView(root *TreeNode) []int {
	arr := make([]int, 0, 100)

	if root == nil {
		return arr
	}
	arr = findRight(root, arr, 0)
	return arr
}

func findRight(head *TreeNode, arr []int, level int) []int {

	if head == nil {
		return arr
	}
	if len(arr) == level {
		arr = append(arr, head.Val)
	}

	arr = findRight(head.Right, arr, level+1)
	arr = findRight(head.Left, arr, level+1)
	return arr
}
