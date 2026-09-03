**365 Days of LeetCode Challenge — Day 4/365**
**Sum of Left Leaves** (Easy)
🔗 https://leetcode.com/problems/sum-of-left-leaves/

Only a parent knows if its child is a "left" leaf. A node can't tell this about itself. So the leaf check happens on `root.Left`, from `root`'s call, not from inside the child's own base case. That flip is the whole trick here.

```go
func sumOfLeftLeaves(root *TreeNode) int {
	if root == nil {
		return 0
	}

	l := sumOfLeftLeaves(root.Left)
	r := sumOfLeftLeaves(root.Right)
	if root.Left != nil && root.Left.Left == nil && root.Left.Right == nil {
		l += root.Left.Val
	}
	return l + r
}
```

Full walkthrough with step-by-step diagrams: https://github.com/architagr/leetcode_solutions/blob/main/easy_problems/401_500/sum_of_left_leaves/SOLUTION.md
