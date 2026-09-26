package countgoodnodesinbinarytree

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// goodNodes seeds the running max with the root's own value, which makes
// the root good by definition. The constraints guarantee a non-nil root.
func goodNodes(root *TreeNode) int {
	return process(root, root.Val)
}

// process counts the good nodes in node's subtree, given the largest value
// on the path from the root down to node's parent.
func process(node *TreeNode, currentMax int) int {
	if node == nil {
		return 0
	}
	res := 0
	// >= because the rule is "nothing on the path is greater": a tie is good.
	if node.Val >= currentMax {
		res++
	}
	// currentMax is a parameter, so this only changes what this node's own
	// children see. A sibling branch keeps the max it was handed.
	currentMax = maxVal(currentMax, node.Val)
	return res + process(node.Left, currentMax) + process(node.Right, currentMax)
}

func maxVal(a, b int) int {
	if a > b {
		return a
	}
	return b
}
