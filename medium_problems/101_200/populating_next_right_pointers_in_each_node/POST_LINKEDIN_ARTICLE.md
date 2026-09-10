---
meta_title: "Reverse the walk and two special cases disappear"
meta_description: "Pushing children right before left makes BFS run each level backwards, so the node to your right is the one you just visited. The assignment becomes trivial."
---

## 365 Days of LeetCode Challenge — Day 50/365

# Populating Next Right Pointers in Each Node

🔗 https://leetcode.com/problems/populating-next-right-pointers-in-each-node/ · Difficulty: Medium

### The problem

Given a perfect binary tree, populate each node's `next` pointer to point at the node to
its right on the same level. The rightmost node of each level points at nil.

![Example 1](images/1.png)

### The intuition

Every node needs a pointer to the node on its right at the same level, and the rightmost
node of each level needs nil. "Same level" makes it a BFS problem, and knowing where a
level ends makes it the nil-sentinel BFS from Day 16.

The obvious way to write it is to walk each level left to right, remembering the previous
node, and set `prev.Next = current`. That works.

This implementation does something neater. It pushes children right before left, so the BFS
visits every level from right to left. And once the walk runs backwards, the node to your
right is simply the node visited just before you — so the assignment becomes
`current.Next = prev`, with no lookahead and nothing written into a node already passed.
The line reads exactly like the requirement, because the traversal was reversed to make
that true.

The level boundary falls out too. When the sentinel is popped, `prev` is set to the
sentinel itself, which is nil. So the first node of the next level — the rightmost one,
since the walk is backwards — gets `Next = nil`, which is precisely what the rightmost node
of a level needs. The rule that would otherwise be a special case is just the general rule
applied at a boundary.

### Builds on

- [Day 16: Binary Tree Zigzag Level Order Traversal](https://github.com/architagr/leetcode_solutions/blob/main/medium_problems/101_200/binary_tree_zigzag_level_order_traversal/) — the nil sentinel in the queue, marking where one level stops and the next begins
- [Day 15: Binary Tree Level Order Traversal](https://github.com/architagr/leetcode_solutions/blob/main/medium_problems/101_200/binary_tree_level_order_traversal/) — level order as the underlying shape

### The solution

```go
func bfs(root *Node) {
	if root == nil {
		return
	}
	// push and pop closures elided
	push(root)
	push(nil)
	var prev *Node
	current := root

	for len(q) > 0 {
		current = pop()
		if current == nil {
			if len(q) > 0 {
				push(nil)
			}
			prev = current
			continue
		}
		current.Next = prev
		prev = current
		if current.Right != nil {
			push(current.Right)
		}
		if current.Left != nil {
			push(current.Left)
		}
	}
}
```

Tracing `[1,2,3,4,5,6,7]`, a perfect tree of three levels.

The root pops first with `prev` still nil, so it gets `Next = nil` — correct, since it's the
only node on its level and therefore the rightmost.

![Step 1: the root gets nil, and children are pushed right-first](images/walkthrough-1.png)

![Step 2: the sentinel resets prev, so the rightmost node gets nil](images/walkthrough-2.png)

Then the rest of the level threads itself: `2` pops after `3`, so `prev` is `3`, which is
indeed the node to its right.

![Step 3: 2 links to 3, the node visited just before it](images/walkthrough-3.png)

![Step 4: the bottom level links the same way](images/walkthrough-4.png)

The re-push guard is the same one as Days 16 and 20 — without `if len(q) > 0`, the final
sentinel would be re-added to an empty queue and the loop would spin forever.

One thing this solution deliberately doesn't use: the tree is guaranteed perfect, and
nothing here depends on that. A perfect tree admits the well-known O(1)-space answer, where
you walk each level using the `Next` pointers already established on the level above,
threading the level below as you go, with no queue at all. This general BFS would work
unchanged on problem 117, where the tree isn't perfect — which is the trade being made.

O(n) time, every node pushed and popped once. Space is O(w) for the queue, where w is the
width of the widest level; in a perfect tree that's roughly n/2, which is precisely what the
O(1) approach exists to avoid.

Full code and the step-by-step walkthrough:
[populating_next_right_pointers_in_each_node](https://github.com/architagr/leetcode_solutions/blob/main/medium_problems/101_200/populating_next_right_pointers_in_each_node/SOLUTION.md)

#DSA #LeetCode #100DaysOfCode #CodingInterview #BinaryTree #BFS #Golang

---

*Solution and code by Archit Agarwal. Write-up drafted with AI assistance from the code and problem statement.*
