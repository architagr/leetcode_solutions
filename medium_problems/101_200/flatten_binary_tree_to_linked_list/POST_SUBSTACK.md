---
meta_title: "Flattening a tree: the order is the whole problem"
meta_description: "The flattened list has to follow pre-order, which means the traversal that produces the answer is one you have already written. Collect, then rebuild."
tags: [golang, binary-tree, recursion, linked-list, leetcode]
---

# Flatten Binary Tree to Linked List

*365 Days of LeetCode Challenge — Day 10/365*

🔗 [LeetCode #114](https://leetcode.com/problems/flatten-binary-tree-to-linked-list/) · Difficulty: Medium

This is the first medium of the binary tree run, and it's a good one to open with because
it doesn't ask you to learn anything new. It asks whether you noticed something on Day 8.

The problem describes an output shape and then, almost in passing, constrains the order of
it. That constraint is the entire problem, and skimming past it is how people end up
writing something much harder than what's being asked.

### The problem

Given the root of a binary tree, flatten it into a "linked list" that reuses the same
`TreeNode` type: every node's left pointer is null, every right pointer points at the next
node, and the order matches a pre-order traversal of the original tree.

![Example 1](images/1.jpg)

### The intuition

The shape being asked for is a linked list wearing `TreeNode`. What makes it a real
problem rather than a pointer exercise is the ordering clause: not sorted, not level by
level, specifically root then left subtree then right subtree.

Which means the traversal that produces the answer is one that's already written.

### Builds on

- [Day 8: Binary Tree Preorder Traversal](https://github.com/architagr/leetcode_solutions/blob/main/easy_problems/101_200/binary_tree_preorder_traversal/) — the exact visit order this problem's output is defined by, and the slice-threading that collects it

Once that lands, the problem splits into two pieces that don't interact. Walk the tree in
pre-order and record the values. Then walk that flat list and rebuild the tree as a
right-leaning chain.

That's a deliberate trade. The follow-up asks for an in-place version using O(1) extra
space, and that one is genuinely fiddly — you're rewiring the pointers you're still
navigating by, so you stash the right subtree before overwriting `Right` with the left
one, then walk to the end of the newly attached run to hang the old right subtree off it.
This solution spends O(n) space and buys code you can read in one pass.

### The solution

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

func getPreOrderArray(root *TreeNode, arr *[]TreeNode) {
	if root == nil {
		return
	}
	(*arr) = append((*arr), TreeNode{
		Val: root.Val,
	})
	getPreOrderArray(root.Left, arr)
	getPreOrderArray(root.Right, arr)
}
```

Tracing `[1,2,5,3,4,null,6]`: node `1` has children `2` and `5`, `2` has children `3` and
`4`, and `5` has a right child `6`. The traversal collects `[1, 2, 3, 4, 5, 6]`, and the
loop chains those into `1 -> 2 -> 3 -> 4 -> 5 -> 6` with every link through `Right`.

Four details that aren't obvious from reading the description.

The traversal appends a `TreeNode` carrying only `Val` — not the node, not a pointer to
it. Nothing from the first phase survives into the second except the numbers, because the
rebuild allocates fresh nodes for everything after the root. The original nodes below the
root become garbage.

The root is the exception. The caller holds that pointer, so it has to remain the same
node; its value gets overwritten from the collected list instead.

`arr[0]` is indexed without a length check, which is safe rather than lucky — the nil
check at the top guarantees the traversal appended at least one value.

And `temp.Left = nil` inside the loop is never load-bearing. On the first pass it
re-clears what the line above already cleared; on every later pass `temp` came from
`new(TreeNode)`, which zeroed it. It stays because it states the loop's invariant
explicitly: nothing this loop leaves behind has a left child.

O(n) time. O(n) space for the collected slice and O(n) for the new chain, plus O(h) for
the recursion stack — exactly what the follow-up's in-place version removes.

---

The habit this one rewards is reading the constraint sentence as carefully as the
requirement sentence. "In the same order as a pre-order traversal" looks like a detail
about the output and is actually the instruction for how to build it.

Full code and the step-by-step walkthrough:
[flatten_binary_tree_to_linked_list](https://github.com/architagr/leetcode_solutions/blob/main/medium_problems/101_200/flatten_binary_tree_to_linked_list/SOLUTION.md)

---

*Solution and code by Archit Agarwal. Write-up drafted with AI assistance from the code and problem statement.*
