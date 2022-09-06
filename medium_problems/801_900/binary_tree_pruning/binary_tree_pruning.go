package binary_tree_pruning

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func PruneTree(root *TreeNode) *TreeNode {
	temp := root
	curr := parse(temp)
	if !curr {
		root = nil
	}
	return root
}
func parse(node *TreeNode) bool {
	if node == nil {
		return false
	}
	curr := node.Val == 1
	right := parse(node.Right)
	if !right {
		node.Right = nil
	}
	left := parse(node.Left)
	if !left {
		node.Left = nil
	}
	return curr || right || left
}
