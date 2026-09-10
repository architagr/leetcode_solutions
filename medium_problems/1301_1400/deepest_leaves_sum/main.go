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
	// The base case is a LEAF, not nil. Only leaves contribute to the
	// answer, and the recursive calls below are guarded, so a nil child is
	// never passed down and never needs representing.
	//
	// The cost: node is dereferenced on the first line without a check,
	// which is safe only because the constraints guarantee at least one
	// node. An empty tree would panic here.
	foo = func(node *TreeNode, level int) {
		if node.Left == nil && node.Right == nil {
			// Deeper than anything seen: everything accumulated so far
			// belonged to a shallower level and is now worthless, so the
			// sum is REPLACED rather than added to.
			//
			// That replacement is why one pass is enough. The final depth
			// never has to be known in advance, because finding a deeper
			// leaf invalidates the previous answer outright.
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
