package diameterofbinarytree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// dia tracks the longest path found so far (in edges), across the whole tree.
// It's package-level so calc's recursive calls can all update the same running max.
var dia = 0

func diameterOfBinaryTree(root *TreeNode) int {
	// Reset before each call since dia is package-level, not local to this call.
	dia = 0
	calc(root)
	return dia
}

// calc returns the height of the subtree rooted at root, and as a side effect
// updates the package-level dia with the best "path through this node" candidate.
func calc(root *TreeNode) int {
	if root == nil {
		// An empty subtree has height 0 and contributes nothing to any path.
		return 0
	}
	left := calc(root.Left)
	right := calc(root.Right)
	// The longest path that passes through root goes left-down then right-down,
	// so its length is left+right edges. Every node gets a turn as the
	// candidate path's center; dia keeps whichever candidate is biggest.
	dia = maxVal(left+right, dia)
	// Height reported to the parent: the taller child subtree, plus one edge
	// down to this node.
	return maxVal(left, right) + 1
}
func maxVal(a, b int) int {
	if a > b {
		return a
	}
	return b
}
