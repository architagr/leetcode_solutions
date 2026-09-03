**365 Days of LeetCode Challenge — Day 17/365**
**Two Sum IV - Input is a BST** (Easy)
🔗 https://leetcode.com/problems/two-sum-iv-input-is-a-bst/

Strip away the "BST" label and this is plain Two Sum. Walk the tree preorder, storing the complement each node still needs (`k - node.Val`), and check every new node against what's already in the map. The map is shared across the whole recursion, so a match can turn up between two totally unrelated branches. That's the part I find neat, and it's also why this never once uses the BST ordering.

```go
func findTarget(root *TreeNode, k int) bool {
	var hashMap map[int]bool = make(map[int]bool)
	return find(root, k, hashMap)
}

func find(root *TreeNode, k int, hashMap map[int]bool) bool {
	if root == nil {
		return false
	}
	if _, ok := hashMap[root.Val]; ok {
		return true
	}
	hashMap[k-root.Val] = true
	left := find(root.Left, k, hashMap)
	right := find(root.Right, k, hashMap)
	return left || right
}
```

Full walkthrough with step-by-step diagrams: https://github.com/architagr/leetcode_solutions/blob/main/easy_problems/601_700/two_sum_iv_input_is_a_bst/SOLUTION.md
