**365 Days of LeetCode Challenge — Day 16/365**
**Find Leaves of Binary Tree** (Medium)
🔗 https://leetcode.com/problems/find-leaves-of-binary-tree/

The problem describes a process: strip the leaves, strip the new leaves, repeat. Simulating that literally is O(n·h).

You don't have to, because a node's round is decided before you remove anything — and it isn't the node's depth. Every per-level problem this week grouped by depth from the root. This one groups by height: the distance down to the deepest leaf beneath a node.

A leaf has nothing below it, so round 0. Everything else waits for both children:

```go
func parse(root *TreeNode, arr *[][]int) int {
	if root == nil {
		return -1
	}
	leftIndex := parse(root.Left, arr)
	rightIndex := parse(root.Right, arr)
	index := maxVal(leftIndex, rightIndex) + 1
	if len(*arr) <= index {
		*arr = append(*arr, []int{})
	}
	(*arr)[index] = append((*arr)[index], root.Val)
	return index
}
```

The `-1` is load-bearing. A leaf's two nil children give `max(-1,-1)+1 = 0`. Return 0 there instead and every round shifts up, leaving an empty first bucket.

O(n) time, one pass.

Full walkthrough with step-by-step diagrams: https://github.com/architagr/leetcode_solutions/blob/main/medium_problems/301_400/find_leaves_of_binary_tree/SOLUTION.md
