package balancedbinarytree

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	_, ok := validateTree(root)
	return ok
}

func validateTree(node *TreeNode) (height int, ok bool) {
	height = 0
	ok = true
	if node == nil {
		return
	}
	rightHeight := 0
	rightHeight, ok = validateTree(node.Right)
	if !ok {
		return
	}
	leftHeight := 0
	leftHeight, ok = validateTree(node.Left)
	if !ok {
		return
	}
	diff := leftHeight - rightHeight
	if diff > 1 || diff < -1 {
		ok = false
		return
	}
	height = maxValue(leftHeight, rightHeight) + 1
	return
}

func maxValue(a, b int) int {
	if a > b {
		return a
	}
	return b
}
