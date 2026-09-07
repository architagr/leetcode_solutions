**365 Days of LeetCode Challenge — Day 18/365**
**Cousins in Binary Tree** (Easy)
🔗 https://leetcode.com/problems/cousins-in-binary-tree/

Cousins means same depth, different parent, so compute both facts for `x` and both for `y` and compare. A plain tree has no parent pointers, but values are unique, so a node's own `Val` can stand in for its parent's identity. The helper checks each child before descending into it, which is the one vantage point where a node's parent is visible.

```go
func isCousins(root *TreeNode, x int, y int) bool {
	depthX, parentX, _ := foo(root, 0, x)
	depthY, parentY, _ := foo(root, 0, y)
	return depthX == depthY && parentX != parentY
}

func foo(node *TreeNode, currentDepth, searchVal int) (depth int, parent int, found bool) {
	depth, parent, found = currentDepth, 0, false
	if node == nil {
		return currentDepth, 0, false
	}
	if node.Val == searchVal {
		return currentDepth, 0, true
	}
	if node.Left != nil {
		if node.Left.Val == searchVal {
			return currentDepth + 1, node.Val, true
		}
		depth, parent, found = foo(node.Left, currentDepth+1, searchVal)
		if found {
			return
		}
	}
	if node.Right != nil {
		if node.Right.Val == searchVal {
			return currentDepth + 1, node.Val, true
		}
		depth, parent, found = foo(node.Right, currentDepth+1, searchVal)
	}
	return
}
```

O(n) time, O(h) space.

Full walkthrough: https://github.com/architagr/leetcode_solutions/blob/main/easy_problems/901_1000/cousins_in_binary_tree/SOLUTION.md
