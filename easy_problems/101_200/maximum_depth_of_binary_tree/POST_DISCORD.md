**365 Days of LeetCode Challenge — Day 1/365**
**Maximum Depth of Binary Tree** (Easy)
🔗 https://leetcode.com/problems/maximum-depth-of-binary-tree/

Count tree levels with BFS. Push a `nil` sentinel right after the root to mark "end of level." Popping it means a level just finished, so bump the depth counter and push a fresh sentinel if there's more tree left to walk. Beats tracking level sizes with a nested loop.

```go
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	max := 0
	queue := make([]*TreeNode, 0)
	queue = append(queue, root, nil)
	pop := func() *TreeNode {
		node := queue[0]
		queue = queue[1:]
		return node
	}
	push := func(node *TreeNode) {
		queue = append(queue, node)
	}
	for len(queue) > 0 {
		node := pop()
		if node == nil {
			if len(queue) > 0 {
				push(nil)
			}
			max++
			continue
		}
		if node.Left != nil {
			push(node.Left)
		}
		if node.Right != nil {
			push(node.Right)
		}
	}

	return max
}
```

Full walkthrough with step-by-step diagrams: https://github.com/architagr/leetcode_solutions/blob/main/easy_problems/101_200/maximum_depth_of_binary_tree/SOLUTION.md
