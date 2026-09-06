package deepestleavessum

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func deepestLeavesSum(root *TreeNode) int {
	deepestLevel := 0
	sumOfDeepestLevel := 0
	var foo func(node *TreeNode, level int)
	foo = func(node *TreeNode, level int) {
		if node.Left == nil && node.Right == nil {
			if level > deepestLevel {
				sumOfDeepestLevel = node.Val
				deepestLevel = level
			} else if level == deepestLevel {
				sumOfDeepestLevel += node.Val
			}
			return
		}
		if node.Left != nil {
			foo(node.Left, level+1)
		}
		if node.Right != nil {
			foo(node.Right, level+1)
		}
	}
	foo(root, 0)
	return sumOfDeepestLevel
}
