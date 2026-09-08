package sumofnodeswithevenvaluedgrandparent

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumEvenGrandparent(root *TreeNode) int {
	return compute(root, nil, nil)
}

func compute(node, parent, grandParent *TreeNode) int {
	if node == nil {
		return 0
	}
	sum := 0
	if grandParent != nil && grandParent.Val%2 == 0 {
		sum += node.Val
	}
	sum += compute(node.Left, node, parent)
	sum += compute(node.Right, node, parent)
	return sum
}
