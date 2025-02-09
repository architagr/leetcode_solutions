package binary_tree_level_order_traversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func LevelOrder(root *TreeNode) [][]int {
	arr := make([][]int, 0)
	return traversal(root, arr, 0)
}

func traversal(head *TreeNode, arr [][]int, level int) [][]int {
	if head == nil {
		return arr
	}
	if (len(arr) - 1) < level {
		a := make([]int, 0)
		arr = append(arr, a)
	}
	arr[level] = append(arr[level], head.Val)
	arr = traversal(head.Left, arr, level+1)
	arr = traversal(head.Right, arr, level+1)
	return arr
}
