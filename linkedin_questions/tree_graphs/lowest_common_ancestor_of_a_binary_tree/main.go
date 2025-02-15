package lowestcommonancestorofabinarytree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	n, _ := f(root, p, q)
	return n
}

func f(root, p, q *TreeNode) (n *TreeNode, count int) {
	if root == nil {
		return nil, 0
	}

	ln, lcount := f(root.Left, p, q)
	if ln != nil {
		n = ln
		return
	}
	rn, rcount := f(root.Right, p, q)
	if rn != nil {
		n = rn
		return
	}
	if root.Val == p.Val || root.Val == q.Val {
		lcount++
	}
	count = lcount + rcount
	if count == 2 {
		n = root
	}
	return
}
