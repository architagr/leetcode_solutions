## 365 Days of LeetCode Challenge — Day 1/365

# Maximum Depth of Binary Tree

🔗 https://leetcode.com/problems/maximum-depth-of-binary-tree/ · Difficulty: Easy

### The problem

Given the root of a binary tree, find its maximum depth: the node count along the
longest path from the root down to the farthest leaf.

### The intuition

Depth here just means level count. Draw the tree on paper and count the rows before
you run out of nodes, and that's your answer.

BFS is the obvious tool for counting rows: visit everything at depth 1, then everything
at depth 2, then depth 3, and keep a tally of how many full rounds you get through
before the queue runs dry.

The annoying part of a plain BFS is that the queue has no idea where one level ends and
the next one starts. You just get a flat stream of nodes with no boundary markers. The
fix I used here is a `nil` sentinel: push one into the queue right after the root.
Popping a `nil` means a level just finished, so bump the depth counter, and if there's
still real work left in the queue, drop in a fresh sentinel to mark the end of the next
one.

That turns "how many levels does this tree have" into "how many sentinels did I pop,"
and a single int handles the bookkeeping. No level-size tracking, no nested `for i := 0,
size := len(queue); i < size; i++` loop like you see in a lot of BFS depth solutions.
Just pop, check if it's nil, and move on.

### The solution

Here's LeetCode's own example tree, which we'll trace below:

![Example tree](tmp-tree.jpg)

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

Tracing it on `[3,9,20,null,null,15,7]` (expected depth `3`):
- Queue starts `[3, nil]`. Pop `3` (real): push its children `9, 20` → queue `[nil, 9, 20]`.
- Pop `nil` (level 1 done, `max=1`): queue not empty, push new sentinel → `[9, 20, nil]`.

  ![Level 1 complete: node 3 visited, max=1](images/walkthrough-1.svg)

- Pop `9` (real, no children) → `[20, nil]`.
- Pop `20` (real): push children `15, 7` → `[nil, 15, 7]`.
- Pop `nil` (level 2 done, `max=2`): push new sentinel → `[15, 7, nil]`.

  ![Level 2 complete: nodes 3, 9, 20 visited, max=2](images/walkthrough-2.svg)

- Pop `15`, then `7` (both leaves) → `[nil]`.
- Pop `nil` (level 3 done, `max=3`): queue empty, no new sentinel.
- Return `max = 3`, which matches.

  ![Level 3 complete: all nodes visited, max=3, loop ends](images/walkthrough-3.svg)

Runtime is O(n): every node gets enqueued and dequeued exactly once. Space is O(n) too
in the worst case, since a very wide, shallow tree can have close to half its nodes
sitting in the queue at the same time.

Full code + tests: `easy_problems/101_200/maximum_depth_of_binary_tree/` in the repo.

#DSA #LeetCode #100DaysOfCode #BinaryTree #BFS #Golang #CodingInterview #Algorithms

---

*Solution and code by Archit Agarwal. Write-up drafted with AI assistance from the code and problem statement.*
