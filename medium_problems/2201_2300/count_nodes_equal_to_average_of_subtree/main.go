package countnodesequaltoaverageofsubtree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func averageOfSubtree(root *TreeNode) int {
	var count = 0
	var sumAndCountOfNodes func(node *TreeNode) (currentNodeSum, countNodes int)
	sumAndCountOfNodes = func(node *TreeNode) (currentNodeSum, countNodes int) {
		currentNodeSum, countNodes = 0, 0
		if node == nil {
			return
		}

		leftSubtreeNodeSum, leftSubTreeNodeCount := sumAndCountOfNodes(node.Left)
		rightSybTreeNodeSum, rightSubTreeNodeCount := sumAndCountOfNodes(node.Right)

		currentNodeSum = leftSubtreeNodeSum + node.Val + rightSybTreeNodeSum
		countNodes = leftSubTreeNodeCount + rightSubTreeNodeCount + 1
		if currentNodeSum/countNodes == node.Val {
			count++
		}
		return
	}
	sumAndCountOfNodes(root)
	return count
}
