package reverseoddlevelsofbinarytree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func reverseOddLevels(root *TreeNode) *TreeNode {
	rev(root.Left, root.Right, 1)
	return root
}

func rev(l, r *TreeNode, d int) {
	if l == nil {
		return
	}
	// At odd levels, swap the mirrored nodes' values
	if d%2 == 1 {
		l.Val, r.Val = r.Val, l.Val
	}
	// Recurse into children, swapping pointers to maintain mirror
	rev(l.Left, r.Right, d+1)
	rev(l.Right, r.Left, d+1)
}
