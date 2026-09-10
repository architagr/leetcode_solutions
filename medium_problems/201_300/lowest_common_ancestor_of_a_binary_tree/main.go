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

// f returns two things: n, the answer once some node has seen both
// targets, and count, how many of the two targets live in this subtree.
//
// The count exists because there is no ordering to exploit here. In the
// BST version (Day 27) a comparison said which way to walk; without one,
// the only way to know whether a target is below a node is to look, so
// subtrees report upward instead of the walk deciding on the way down.
func f(root, p, q *TreeNode) (n *TreeNode, count int) {
	if root == nil {
		return nil, 0
	}

	ln, lcount := f(root.Left, p, q)
	// Early return is what keeps the answer LOWEST. Without it the walk
	// continues and every ancestor above the real LCA also reaches a
	// count of 2, overwriting the answer with itself.
	if ln != nil {
		n = ln
		return
	}
	rn, rcount := f(root.Right, p, q)
	if rn != nil {
		n = rn
		return
	}
	// Folded into lcount rather than a separate variable: a node counts
	// toward its own subtree's total, which is what makes "a node may be
	// a descendant of itself" work. If p is an ancestor of q, then at p
	// the count reaches 2 - one for being p, one from q's subtree.
	//
	// Compared on Val, which relies on the problem's guarantee that
	// values are unique; with duplicates this would need pointers.
	if root.Val == p.Val || root.Val == q.Val {
		lcount++
	}
	count = lcount + rcount
	// Post-order means nodes are reached bottom-up, so the first node to
	// reach 2 is the lowest one that can.
	if count == 2 {
		n = root
	}
	return
}
