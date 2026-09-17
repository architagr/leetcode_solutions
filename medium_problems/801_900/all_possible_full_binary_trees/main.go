package sum_of_distances_in_tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// memo caches, per node count, every full binary tree of that size. A subtree
// of k nodes looks the same wherever it hangs, so it is only enumerated once.
var memo map[int][]*TreeNode

func allPossibleFBT(n int) []*TreeNode {
	memo = make(map[int][]*TreeNode)
	return build(n)
}

// build returns every full binary tree holding exactly n nodes.
func build(n int) []*TreeNode {
	// A full binary tree has a root plus two subtrees of equal parity, so its
	// node count is always odd.
	if n%2 == 0 {
		return nil
	}
	if n == 1 {
		return []*TreeNode{{Val: 0}}
	}
	if cached, ok := memo[n]; ok {
		return cached
	}

	res := make([]*TreeNode, 0)
	// The root eats one node. Every odd split of the remaining n-1 between the
	// two subtrees is a distinct family of trees, and each left shape pairs with
	// each right shape.
	for left := 1; left <= n-2; left += 2 {
		right := n - 1 - left
		for _, l := range build(left) {
			for _, r := range build(right) {
				res = append(res, &TreeNode{
					Val:   0,
					Left:  copyBinaryTree(l),
					Right: copyBinaryTree(r),
				})
			}
		}
	}

	memo[n] = res
	return res
}

func copyBinaryTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	return &TreeNode{
		Val:   root.Val,
		Left:  copyBinaryTree(root.Left),
		Right: copyBinaryTree(root.Right),
	}
}
