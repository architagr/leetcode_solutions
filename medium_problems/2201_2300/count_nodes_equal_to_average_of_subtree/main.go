package countnodesequaltoaverageofsubtree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func averageOfSubtree(root *TreeNode) int {
	var count = 0
	// Returns (sum of subtree, number of nodes in subtree)
	var sumAndCountOfNodes func(node *TreeNode) (currentNodeSum, countNodes int)
	sumAndCountOfNodes = func(node *TreeNode) (currentNodeSum, countNodes int) {
		currentNodeSum, countNodes = 0, 0
		if node == nil {
			return
		}

		// Recursively compute sums and counts for both subtrees
		leftSubtreeNodeSum, leftSubTreeNodeCount := sumAndCountOfNodes(node.Left)
		rightSybTreeNodeSum, rightSubTreeNodeCount := sumAndCountOfNodes(node.Right)

		// Combine subtree values with current node
		currentNodeSum = leftSubtreeNodeSum + node.Val + rightSybTreeNodeSum
		countNodes = leftSubTreeNodeCount + rightSubTreeNodeCount + 1
		// Check if node value equals subtree average
		if currentNodeSum/countNodes == node.Val {
			count++
		}
		return
	}
	sumAndCountOfNodes(root)
	return count
}
