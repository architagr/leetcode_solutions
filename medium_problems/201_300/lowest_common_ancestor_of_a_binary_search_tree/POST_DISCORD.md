**365 Days of LeetCode Challenge — Day 27/365**
**Lowest Common Ancestor of a Binary Search Tree** (Medium)
🔗 https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-search-tree/

The general version of this needs a search. The BST version doesn't, for the same reason Day 26 didn't: a comparison in a BST is a direction, not a verdict.

Run it for both targets at once and there are three cases. Both smaller than the node — both live left, so the answer is further left. Both larger — further right. Anything else is the answer.

That last line is doing more than it looks. It covers targets on opposite sides (this node is where their paths diverge) *and* one target being the node itself (the LCA definition lets a node be its own descendant). Nothing has to detect which — "not both smaller, not both larger" is exactly the union.

```go
if root.Val > p.Val && root.Val > q.Val {
	return lowestCommonAncestor(root.Left, p, q)
} else if root.Val < p.Val && root.Val < q.Val {
	return lowestCommonAncestor(root.Right, p, q)
}
return root
```

There's no search for p or q anywhere — the problem guarantees they exist, so the walk only has to find where they stop agreeing on a direction.

O(h) time. Every call is in tail position, so a `for` loop version would be O(1) space.

Full walkthrough with step-by-step diagrams: https://github.com/architagr/leetcode_solutions/blob/main/medium_problems/201_300/lowest_common_ancestor_of_a_binary_search_tree/SOLUTION.md
