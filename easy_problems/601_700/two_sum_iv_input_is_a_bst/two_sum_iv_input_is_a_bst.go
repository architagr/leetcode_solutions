package two_sum_iv_input_is_a_bst

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findTarget(root *TreeNode, k int) bool {
	// One map shared across the whole recursion (maps are reference types
	// in Go), so a complement stored while visiting one branch is still
	// visible when a completely different branch is visited later.
	var hashMap map[int]bool = make(map[int]bool)

	return find(root, k, hashMap)
}

// find walks the tree preorder (this node, then left subtree, then right
// subtree). It's the classic single-pass Two Sum hash-set trick ported onto
// a tree traversal instead of an array iteration -- it never relies on the
// BST ordering, so it would work on any binary tree.
func find(root *TreeNode, k int, hashMap map[int]bool) bool {
	if root == nil {
		// Empty subtree can't contain a matching pair.
		return false
	}
	if _, ok := hashMap[root.Val]; ok {
		// Some earlier-visited node already recorded root.Val as the
		// complement it needed (i.e. that node's value + root.Val == k).
		// Return immediately without recursing into this node's subtree.
		return true
	}
	// No match yet: remember what value this node would need to see later
	// in order to complete a pair.
	hashMap[k-root.Val] = true
	// Note: left and right are separate statements, not a single
	// short-circuited `find(...) || find(...)` expression, so right is
	// always evaluated even when left already found a match.
	left := find(root.Left, k, hashMap)
	right := find(root.Right, k, hashMap)
	return left || right
}
