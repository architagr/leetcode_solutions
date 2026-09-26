package sumofnodeswithevenvaluedgrandparent

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumEvenGrandparent(root *TreeNode) int {
	// The root has neither a parent nor a grandparent; nil stands for "none".
	return compute(root, nil, nil)
}

// compute sums the values in node's subtree whose grandparent is even.
// Nodes have no parent pointers, so the recursion carries the last two
// nodes on the path down instead.
func compute(node, parent, grandParent *TreeNode) int {
	if node == nil {
		return 0
	}
	sum := 0
	// The nil test short-circuits, so .Val is never read on a missing
	// grandparent (the root and its children).
	if grandParent != nil && grandParent.Val%2 == 0 {
		sum += node.Val
	}
	// Shift the window down one level: the child's parent is this node, and
	// its grandparent is this node's parent.
	sum += compute(node.Left, node, parent)
	sum += compute(node.Right, node, parent)
	return sum
}
