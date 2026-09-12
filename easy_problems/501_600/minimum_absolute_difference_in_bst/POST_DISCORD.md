**365 Days of LeetCode Challenge — Day 47/365**
**Minimum Absolute Difference in BST** (Easy)
🔗 https://leetcode.com/problems/minimum-absolute-difference-in-bst/

In-order traversal of a BST visits nodes in strictly increasing order, and in any sorted sequence the smallest gap always sits between neighbors. So instead of checking every pair (O(n²)), compare each node to the previous node hit during the traversal and keep the smallest gap. Sorting for free, basically.

```go
func getMinimumDifference(root *TreeNode) int {
	res := math.MaxInt
	var prev *TreeNode
	var helper func(root *TreeNode)
	helper = func(root *TreeNode) {
		if root == nil {
			return
		}

		helper(root.Left)

		if prev != nil {
			res = min(res, root.Val-prev.Val)
		}
		prev = root

		helper(root.Right)
	}

	helper(root)

	return res
}
```

Since traversal order is sorted, `root.Val >= prev.Val` always holds, so no abs() is needed. O(n) time, O(h) space for the recursion stack.

Full walkthrough with step-by-step diagrams: https://github.com/architagr/leetcode_solutions/blob/main/easy_problems/501_600/minimum_absolute_difference_in_bst/SOLUTION.md
