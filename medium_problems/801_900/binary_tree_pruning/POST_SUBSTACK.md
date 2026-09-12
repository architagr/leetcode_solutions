---
meta_title: "Binary Tree Pruning: the parent does the deleting"
meta_description: "A node cannot remove itself, having no parent pointer. So each child reports whether it holds a 1, and the parent clears the pointer it owns."
tags: [golang, binary-tree, recursion, postorder, leetcode]
---

# Binary Tree Pruning

*365 Days of LeetCode Challenge — Day 71/365*

🔗 [LeetCode #814](https://leetcode.com/problems/binary-tree-pruning/) · Difficulty: Medium

The condition in this problem is short enough to guess correctly on the first read. The
mechanics of applying it are where the actual work is, and they come down to a limitation of
the data structure rather than anything about the question.

A binary tree node points at its children and nothing else. That single fact decides the
shape of every solution that removes nodes from one.

### The problem

Return the same tree with every subtree not containing a `1` removed.

![Example 1](images/1.png)

### The intuition

A node survives if its own subtree contains a `1` anywhere. Written as a recursion, that's
almost the definition read aloud: this node keeps its place if it is a `1`, or if either of
its subtrees keeps anything.

The direction matters. A node cannot answer that question on the way down — it doesn't yet
know what's beneath it. So this is post-order: both children report first, and the node
combines their answers with its own value.

The part that's easy to get wrong is who does the deleting. A node can't remove itself — it
has no reference to its own parent, and nulling a local variable changes nothing the caller
can see. So the pruning is done by the parent, immediately after each recursive call
returns. The child reports "nothing worth keeping down here", and the parent is the one
holding the pointer that has to be cleared.

Which leaves the root, because the root has no parent. That's why `PruneTree` exists as a
wrapper at all: it calls `parse`, and if the whole tree reports back false, it sets `root`
to nil itself. Without that, a tree of all zeroes would come back intact instead of empty.

### Builds on

- [Day 9: Binary Tree Postorder Traversal](https://github.com/architagr/leetcode_solutions/blob/main/easy_problems/101_200/binary_tree_postorder_traversal/) — the children-before-parent order that makes a bottom-up answer possible
- [Day 69: Evaluate Boolean Binary Tree](https://github.com/architagr/leetcode_solutions/blob/main/easy_problems/2301_2400/evaluate_boolean_binary_tree/) — the same shape, where each node returned a boolean built from its children's booleans

### The solution

```go
func PruneTree(root *TreeNode) *TreeNode {
	temp := root
	curr := parse(temp)
	if !curr {
		root = nil
	}
	return root
}

func parse(node *TreeNode) bool {
	if node == nil {
		return false
	}
	curr := node.Val == 1
	right := parse(node.Right)
	if !right {
		node.Right = nil
	}
	left := parse(node.Left)
	if !left {
		node.Left = nil
	}
	return curr || right || left
}
```

Tracing `[1,null,0,0,1]`: root `1` with no left child and a right child `0`, whose children
are `0` and `1`. Expected output `[1,null,0,null,1]`.

![Step 1: a leaf 0 reports false](images/walkthrough-1.png)

![Step 2: the parent clears the pointer to the pruned child](images/walkthrough-2.png)

Node `0` is not a `1`, but its right subtree kept the leaf `1`, so it survives — which is
the case that shows why "prune every 0" would be the wrong rule.

![Step 3: a 0 survives because its subtree kept a 1](images/walkthrough-3.png)

![Step 4: the root reports true and is kept](images/walkthrough-4.png)

The nil base case returning `false` does quiet work: an absent subtree contains no `1`, so
it contributes nothing to the `||`, and the parent's `node.Right = nil` becomes a harmless
no-op on a pointer that was already nil.

`parse` recurses right before left, and here that's arbitrary — each subtree's answer is
independent of the other's.

One note on the shape. This returns a boolean and mutates the tree as a side effect. The
more common formulation returns `*TreeNode` and has the caller reassign — `root.Left =
prune(root.Left)` — which is what Day 56 used for deletion. Both work, and the reassigning
version needs no wrapper, because returning nil for the root handles the whole-tree case
naturally.

O(n) time, every node visited once. Space is O(h) for the recursion stack.

---

Any operation that removes nodes from a singly-linked structure runs into the same wall: the
thing being removed is not the thing holding the reference. The fix is always to move the
decision up one level, and then to notice that the top level has no one above it.

That is what the wrapper function is for, and it is why a solution can look complete, pass
most tests, and still fail on an input where the whole tree should disappear.

Full code and the step-by-step walkthrough:
[binary_tree_pruning](https://github.com/architagr/leetcode_solutions/blob/main/medium_problems/801_900/binary_tree_pruning/SOLUTION.md)

---

*Solution and code by Archit Agarwal. Write-up drafted with AI assistance from the code and problem statement.*
