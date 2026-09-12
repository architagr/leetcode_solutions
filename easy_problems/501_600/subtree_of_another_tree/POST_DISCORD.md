**365 Days of LeetCode Challenge — Day 68/365**
**Subtree of Another Tree** (Easy)
🔗 https://leetcode.com/problems/subtree-of-another-tree/

This is "are two trees identical" (the classic Same Tree check), run at every node of `root` as a candidate anchor. Checking values before the full structural comparison saves a lot of wasted recursion, and honestly that one comparison is doing most of the work here.

```go
func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	if subRoot == nil {
		return true
	}
	if root == nil {
		return subRoot == nil
	}

	if root.Val == subRoot.Val && equalBinaryTree(root, subRoot) {
		return true
	}
	return isSubtree(root.Left, subRoot) || isSubtree(root.Right, subRoot)
}

func equalBinaryTree(root *TreeNode, subRoot *TreeNode) bool {
	if root == nil {
		return subRoot == nil
	}
	if subRoot == nil {
		return root == nil
	}
	if root.Val != subRoot.Val {
		return false
	}
	return equalBinaryTree(root.Left, subRoot.Left) && equalBinaryTree(root.Right, subRoot.Right)
}
```

O(m·n) time (`m` = nodes in root, `n` = nodes in subRoot), O(h1 + h2) space for the two recursion stacks.

Full walkthrough with step-by-step diagrams: https://github.com/architagr/leetcode_solutions/blob/main/easy_problems/501_600/subtree_of_another_tree/SOLUTION.md
