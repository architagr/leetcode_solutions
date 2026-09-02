package leafsimilartrees

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	// Extract each tree's leaf sequence independently; from here on we only
	// ever compare two int slices, not the trees themselves.
	root1Leafs := leafs(root1)
	root2Leafs := leafs(root2)
	// Cheap early exit: different leaf counts can never be leaf-similar.
	if len(root1Leafs) != len(root2Leafs) {
		return false
	}
	for i := 0; i < len(root1Leafs); i++ {
		if root1Leafs[i] != root2Leafs[i] {
			return false
		}
	}
	return true
}

// leafs returns node's leaves in left-to-right order.
func leafs(node *TreeNode) []int {
	if node == nil {
		return []int{}
	}
	// A childless node is a leaf; contribute just its own value.
	if node.Left == nil && node.Right == nil {
		return []int{node.Val}
	}
	// Left subtree's leaves always come before the right subtree's, so the
	// concatenation below naturally preserves left-to-right order.
	return append(leafs(node.Left), leafs(node.Right)...)
}
