package evaluatebooleanbinarytree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

const (
	FALSE = 0
	TRUE  = 1
	OR    = 2
	AND   = 3
)

func evaluateTree(root *TreeNode) bool {
	if root.Left == nil && root.Right == nil {
		return root.Val == TRUE
	}

	left := evaluateTree(root.Left)
	right := evaluateTree(root.Right)
	if root.Val == OR {
		return left || right
	}
	return left && right
}
