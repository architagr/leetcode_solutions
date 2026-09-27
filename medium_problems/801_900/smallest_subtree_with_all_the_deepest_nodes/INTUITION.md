## Intuition

"The smallest subtree containing all the deepest nodes" is a long way of saying "the lowest
common ancestor of the deepest leaves". A subtree contains a node exactly when its root is an
ancestor of that node, and the smallest such subtree hangs from the lowest shared ancestor.

This solution works it out from paths. A DFS carries the root-to-node path. At each leaf, if
the path is as long as the deepest seen so far, a copy is saved (shorter ones are dropped as
deeper leaves appear). Afterwards every saved path starts at the root and runs to a deepest
leaf. Walk them in parallel from the root: the last position where they all hold the same node
is the answer.

## Builds on

- [Day 49: Lowest Common Ancestor of a Binary Tree](https://github.com/architagr/leetcode_solutions/blob/main/medium_problems/201_300/lowest_common_ancestor_of_a_binary_tree/) — the smallest subtree containing a set of nodes is their lowest common ancestor
- [Day 34: Deepest Leaves Sum](https://github.com/architagr/leetcode_solutions/blob/main/medium_problems/1301_1400/deepest_leaves_sum/) — finding which leaves are the deepest

It's correct and easy to follow, but it stores a copy of every deepest path, which can be
O(n · h) memory, and it keeps its state in package-level variables (`data`, `maxDepth`), so
two calls at once would interfere.

A single postorder pass does the same job in O(n) time without any paths. Each call returns
the depth of its subtree and the answer within it. If the left side is deeper, the deepest
leaves are all on the left, so pass up the left answer. Same for the right. If both sides are
equally deep, deepest leaves exist on both sides, so this node is where they meet:

```go
func deepestSubtree(root *TreeNode) *TreeNode {
	var walk func(n *TreeNode) (int, *TreeNode)
	walk = func(n *TreeNode) (int, *TreeNode) {
		if n == nil {
			return 0, nil
		}
		ld, l := walk(n.Left)
		rd, r := walk(n.Right)
		switch {
		case ld > rd:
			return ld + 1, l
		case rd > ld:
			return rd + 1, r
		default:
			return ld + 1, n
		}
	}
	_, node := walk(root)
	return node
}
```

I checked it returns the very same node as the path version on three thousand random trees.
Tomorrow's problem is this one again (LeetCode lists them as duplicates), with a different path
technique in the repo.

**Complexity (as written):**
- Time: O(n · h) to copy and compare paths.
- Space: O(n · h) for the saved paths in the worst case.
