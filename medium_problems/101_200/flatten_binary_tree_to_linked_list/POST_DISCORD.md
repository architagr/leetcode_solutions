**365 Days of LeetCode Challenge — Day 10/365**
**Flatten Binary Tree to Linked List** (Medium)
🔗 https://leetcode.com/problems/flatten-binary-tree-to-linked-list/

The shape wanted is a linked list wearing `TreeNode`. The clause that makes it a real problem is the ordering — not sorted, not level by level, specifically pre-order. So the traversal that produces the answer is Day 8's.

That splits it into two halves that don't interact: collect the values in pre-order, then rebuild as a right-leaning chain. The follow-up asks for in-place with O(1) space, which is a genuinely different problem — you'd be rewiring the pointers you're still navigating by.

```go
func Flatten(root *TreeNode) {
	if root == nil {
		return
	}
	temp := root
	arr := make([]TreeNode, 0)
	getPreOrderArray(temp, &arr)

	root.Val = arr[0].Val
	temp = root
	temp.Left = nil

	for i := 1; i < len(arr); i++ {
		y := new(TreeNode)
		y.Val = arr[i].Val
		temp.Right = y
		temp.Left = nil
		temp = temp.Right
	}
}
```

O(n) time, O(n) space for the collected slice plus the new chain.

Full walkthrough with step-by-step diagrams: https://github.com/architagr/leetcode_solutions/blob/main/medium_problems/101_200/flatten_binary_tree_to_linked_list/SOLUTION.md
