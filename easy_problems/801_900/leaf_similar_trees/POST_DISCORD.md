**365 Days of LeetCode Challenge — Day 22/365**
**Leaf-Similar Trees** (Easy)
🔗 https://leetcode.com/problems/leaf-similar-trees/

Forget the tree shapes entirely. Pull each tree's leaves out left-to-right into a slice (the left subtree always finishes before the right one starts, so the order comes out correct on its own). After that it's just list equality: same length, same value at every index.

```go
func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	root1Leafs := leafs(root1)
	root2Leafs := leafs(root2)
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

func leafs(node *TreeNode) []int {
	if node == nil {
		return []int{}
	}
	if node.Left == nil && node.Right == nil {
		return []int{node.Val}
	}
	return append(leafs(node.Left), leafs(node.Right)...)
}
```

O(n + m) time, O(n + m) space plus the recursion stack.

Full walkthrough with step-by-step diagrams: https://github.com/architagr/leetcode_solutions/blob/main/easy_problems/801_900/leaf_similar_trees/SOLUTION.md
