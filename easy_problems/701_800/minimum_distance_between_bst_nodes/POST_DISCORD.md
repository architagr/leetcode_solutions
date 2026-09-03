**365 Days of LeetCode Challenge — Day 21/365**
**Minimum Distance Between BST Nodes** (Easy)
🔗 https://leetcode.com/problems/minimum-distance-between-bst-nodes/

An in-order traversal of a BST hands back the values already sorted, so the smallest difference between any two nodes can only sit between two neighbours in that list. Traverse once, then scan adjacent pairs. The whole solve is just trusting that the tree sorts itself for you.

```go
func minDiffInBST(root *TreeNode) int {
	if root == nil || (root.Left == nil && root.Right == nil) {
		return 0
	}
	arr := inOrder(root)
	minVal := math.MaxInt
	for i := 1; i < len(arr); i++ {
		minVal = min(minVal, abs(arr[i]-arr[i-1]))
	}
	return minVal
}

func inOrder(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	return append(inOrder(root.Left), append([]int{root.Val}, inOrder(root.Right)...)...)
}
```

Full walkthrough with step-by-step diagrams: https://github.com/architagr/leetcode_solutions/blob/main/easy_problems/701_800/minimum_distance_between_bst_nodes/SOLUTION.md
