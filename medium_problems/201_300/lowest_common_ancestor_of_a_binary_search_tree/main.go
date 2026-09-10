package lowestcommonancestorofabinarysearchtree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	// Degenerate input only. The problem guarantees both nodes exist, so
	// neither branch fires on a valid call: two nils means every node
	// trivially qualifies, one nil means there is no sensible answer.
	if p == nil && q == nil {
		return root
	} else if p == nil || q == nil {
		return nil
	}
	// Both targets smaller: in a BST they can only be in the left subtree,
	// so the node they share is there too and cannot be this one. One
	// comparison per target eliminates a whole side, with no search and no
	// backtracking.
	if root.Val > p.Val && root.Val > q.Val {
		return lowestCommonAncestor(root.Left, p, q)
	} else if root.Val < p.Val && root.Val < q.Val {
		return lowestCommonAncestor(root.Right, p, q)
	}
	// Everything else is the answer, and this one line covers two cases:
	// the targets straddle this node, so the paths to them diverge here;
	// or one target IS this node, which counts because the LCA definition
	// lets a node be a descendant of itself. "Not both smaller, not both
	// larger" is exactly the union, so neither needs detecting.
	return root
}
