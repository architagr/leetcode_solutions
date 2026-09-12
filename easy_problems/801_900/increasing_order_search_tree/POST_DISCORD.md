**365 Days of LeetCode Challenge — Day 53/365**
**Increasing Order Search Tree** (Easy)
🔗 https://leetcode.com/problems/increasing-order-search-tree/

The target shape (no left child, one right child everywhere) is just a sorted linked list built out of `TreeNode`s, and an in-order walk of a BST already visits nodes in ascending order. So: collect in-order, then relink with `.Right`. The part worth calling out is clearing each node's old child pointers as it's collected, otherwise stale links from the original tree survive into the final vine.

```go
func increasingBST(root *TreeNode) *TreeNode {
	inorder := inOrder(root)
	for i := 1; i < len(inorder); i++ {
		inorder[i-1].Right = inorder[i]
	}
	return inorder[0]
}

func inOrder(root *TreeNode) []*TreeNode {
	if root == nil {
		return []*TreeNode{}
	}
	res := make([]*TreeNode, 0)
	left := inOrder(root.Left)
	right := inOrder(root.Right)

	root.Left = nil
	root.Right = nil
	res = append(left, root)
	res = append(res, right...)
	return res
}
```

O(n) time, O(n) space for the list plus O(h) recursion stack. No new nodes get allocated.

Full walkthrough with step-by-step diagrams: https://github.com/architagr/leetcode_solutions/blob/main/easy_problems/801_900/increasing_order_search_tree/SOLUTION.md
