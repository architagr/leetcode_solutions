package symmetrictree

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return isSameTree(root.Left, invertTree(root.Right))
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	} else if p == nil || q == nil {
		return false
	}
	if p.Val != q.Val {
		return false
	}
	left := isSameTree(p.Left, q.Left)
	right := isSameTree(p.Right, q.Right)

	return left && right

}

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	root.Left = invertTree(root.Left)
	root.Right = invertTree(root.Right)
	root.Right, root.Left = root.Left, root.Right
	return root
}
