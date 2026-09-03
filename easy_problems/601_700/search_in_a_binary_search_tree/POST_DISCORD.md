**365 Days of LeetCode Challenge — Day 19/365**
**Search in a Binary Search Tree** (Easy)
🔗 https://leetcode.com/problems/search-in-a-binary-search-tree/

At each node, comparing `root.Val` against `val` tells you which single subtree could still hold the answer, so you never check both children and never backtrack. It's a straight walk from the root down to wherever `val` lives, or down to `nil` if it isn't there. One comparison per level, one subtree gone per comparison.

```go
func searchBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val > val {
		return searchBST(root.Left, val)
	}
	if root.Val < val {
		return searchBST(root.Right, val)
	}
	return root
}
```

Full walkthrough with step-by-step diagrams: https://github.com/architagr/leetcode_solutions/blob/main/easy_problems/601_700/search_in_a_binary_search_tree/SOLUTION.md
