**365 Days of LeetCode Challenge — Day 31/365**
**Binary Tree Right Side View** (Medium)
🔗 https://leetcode.com/problems/binary-tree-right-side-view/

Standing on the right and seeing what isn't hidden has a plainer description: from each level you see exactly one node, the rightmost.

That sounds like BFS, and BFS works. This one is a DFS with two small changes to the Day 29 shape. Store only the first node reached at each depth — `arr` holds one value per level, so `len(arr) == level` means "first arrival here". Then visit right before left, which makes that first arrival the rightmost node.

Neither piece works alone. The guard picks first arrivals; the traversal order decides which arrival is first. Swap the two recursive calls and you get the left side view.

```go
func findRight(head *TreeNode, arr []int, level int) []int {
	if head == nil {
		return arr
	}
	if len(arr) == level {
		arr = append(arr, head.Val)
	}
	arr = findRight(head.Right, arr, level+1)
	arr = findRight(head.Left, arr, level+1)
	return arr
}
```

The left subtree still gets walked, because the rightmost node at a level isn't always in the right subtree — that branch can just run out of depth first.

O(n) time, O(h) space.

Full walkthrough with step-by-step diagrams: https://github.com/architagr/leetcode_solutions/blob/main/medium_problems/101_200/binary_tree_right_side_view/SOLUTION.md
