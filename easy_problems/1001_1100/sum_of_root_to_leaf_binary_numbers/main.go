package sumofroottoleafbinarynumbers

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumRootToLeaf(root *TreeNode) int {
	// rootLevelSum starts at 0: no bits accumulated yet above the root.
	return sum(root, 0)
}

// sum returns the total of all root-to-leaf binary numbers in this subtree.
// rootLevelSum is the binary number built from the true root down to (but
// not including) this node, i.e. the bits contributed by ancestors so far.
func sum(root *TreeNode, rootLevelSum int) int {
	if root == nil {
		return 0
	}
	// Shift the accumulated bits one place left to make room for this
	// node's bit, then drop root.Val (0 or 1) into the new units place —
	// the same way a binary number grows digit by digit left to right.
	rootLevelSum *= 2
	currentSum := rootLevelSum + root.Val
	if isLeafNode(root) {
		// currentSum already is the complete binary number for this
		// root-to-leaf path, so it's the whole contribution from here.
		return currentSum
	}

	// Not a leaf: pass the extended currentSum down as the new
	// rootLevelSum for each child, and sum whatever they each contribute.
	leftSum := sum(root.Left, currentSum)
	rightSum := sum(root.Right, currentSum)
	return leftSum + rightSum
}

func isLeafNode(node *TreeNode) bool {
	return node.Left == nil && node.Right == nil
}
