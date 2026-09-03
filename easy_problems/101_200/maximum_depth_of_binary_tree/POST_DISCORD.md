**365 Days of LeetCode Challenge — Day 1/365**
**Maximum Depth of Binary Tree** (Easy)
🔗 https://leetcode.com/problems/maximum-depth-of-binary-tree/

**Intuition:** count tree levels with BFS. Push a `nil` sentinel right after the root
to mark "end of level." Popping it means a level just finished, so bump the depth
counter and push a fresh sentinel if there's more tree left to walk. Beats tracking
level sizes with a nested loop.

Here's LeetCode's own example tree, which we'll trace below:

![Example tree](tmp-tree.jpg)

**Full solution:**
```go
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	max := 0
	queue := make([]*TreeNode, 0)
	// nil is a sentinel marking "end of the current level" in the queue, so a
	// plain BFS can tell levels apart without tracking per-level sizes.
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
			// Popping the sentinel means we've drained a whole level.
			if len(queue) > 0 {
				// More real nodes remain: push the next level's end marker.
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

**Walkthrough** on `[3,9,20,null,null,15,7]` (expected `3`):
- `[3,nil]` → pop `3`, push children → `[nil,9,20]`
- pop `nil` → level 1 done (`max=1`), push new sentinel → `[9,20,nil]`

![Level 1 complete: node 3 visited, max=1](images/walkthrough-1.svg)

- pop `9` (leaf) → `[20,nil]`
- pop `20`, push children → `[nil,15,7]`
- pop `nil` → level 2 done (`max=2`), push sentinel → `[15,7,nil]`

![Level 2 complete: nodes 3, 9, 20 visited, max=2](images/walkthrough-2.svg)

- pop `15`, `7` (leaves) → `[nil]`
- pop `nil` → level 3 done (`max=3`), queue empty, stop
- return 3, matches

![Level 3 complete: all nodes visited, max=3, loop ends](images/walkthrough-3.svg)

O(n) time, O(n) worst-case space for a wide, shallow tree where the queue can hold
close to half the nodes at once.
