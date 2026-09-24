package countgoodnodesinbinarytree

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func goodNodes(root *TreeNode) int {
	return process(root, root.Val)
}

func process(node *TreeNode, currentMax int) int {
	if node == nil {
		return 0
	}
	res := 0
	if node.Val >= currentMax {
		res++
	}
	currentMax = maxVal(currentMax, node.Val)
	return res + process(node.Left, currentMax) + process(node.Right, currentMax)
}

func maxVal(a, b int) int {
	if a > b {
		return a
	}
	return b
}
