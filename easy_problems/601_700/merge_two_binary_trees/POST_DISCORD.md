**365 Days of LeetCode Challenge — Day 40/365**
**Merge Two Binary Trees** (Easy)
🔗 https://leetcode.com/problems/merge-two-binary-trees/

Sum the values where both trees have a node, and everywhere else just hand back whichever subtree is actually there. Nothing gets built for the parts that don't overlap, so a new node only ever appears where both sides genuinely meet.

```go
func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	if root1 == nil {
		return root2
	}
	if root2 == nil {
		return root1
	}
	root := &TreeNode{Val: root1.Val + root2.Val}
	root.Left = mergeTrees(root1.Left, root2.Left)
	root.Right = mergeTrees(root1.Right, root2.Right)
	return root
}
```

Full walkthrough with step-by-step diagrams: https://github.com/architagr/leetcode_solutions/blob/main/easy_problems/601_700/merge_two_binary_trees/SOLUTION.md
