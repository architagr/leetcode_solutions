**365 Days of LeetCode Challenge — Day 49/365**
**Lowest Common Ancestor of a Binary Tree** (Medium)
🔗 https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-tree/

Yesterday's problem with one guarantee removed, and it's worth seeing how much that guarantee was doing.

In a BST, comparing a target against a node told you which subtree it was in, so the walk went down one path and ignored the rest of the tree. Here there's no ordering, so a comparison tells you nothing about direction. The only way to know if a target is under a node is to go and look.

That flips the shape. Yesterday walked downward deciding as it went; this walks everything and reports upward. Each subtree answers "how many of the two targets are in me?", the parent adds them, and the first node to reach 2 is the LCA — post-order reaches nodes bottom-up, so the first one to hit 2 is the lowest that can.

```go
ln, lcount := f(root.Left, p, q)
if ln != nil {
	n = ln
	return
}
```

That early return is what keeps it *lowest*. Without it every ancestor above also reaches 2 and overwrites the answer with itself.

O(n) time versus yesterday's O(h). That gap is the cost of losing the ordering.

Full walkthrough with step-by-step diagrams: https://github.com/architagr/leetcode_solutions/blob/main/medium_problems/201_300/lowest_common_ancestor_of_a_binary_tree/SOLUTION.md
