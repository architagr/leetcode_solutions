## 365 Days of LeetCode Challenge — Day 27/365

# Find a Corresponding Node of a Binary Tree in a Clone of That Tree

🔗 https://leetcode.com/problems/find-a-corresponding-node-of-a-binary-tree-in-a-clone-of-that-tree/ · Difficulty: Easy

### The problem

You get two binary trees, `original` and `cloned`. `cloned` is guaranteed to be an exact
structural copy: same shape, same values, just different node objects living in memory.
You're also handed `target`, a reference to one specific node inside `original`. The job
is to hand back the matching node inside `cloned`, without touching either tree or
`target` itself.

### The intuition

The detail that makes this one easy is buried in the setup and easy to skim past:
`cloned` is a structurally identical copy of `original`. Same shape, same values, node
for node. Once that lands, you stop thinking about searching `cloned` on its own and
trying to recognize where `target` sits inside it. Instead you walk both trees at the
same time, one step down `original` and the matching step down `cloned` inside a single
recursive call, and the two pointers are always looking at the same logical spot in the
tree, just in two different objects.

At each paired position the question is just: is this the node? Compare values. If they
match, the `cloned`-side pointer is sitting exactly where `target` sits in `original`, so
that's the answer and the search stops right there.

If it's not a match, recurse left in both trees together first. Anything that search
finds gets passed straight back up the call stack. If it comes back empty, recurse right
in both trees together and return whatever that finds instead. Because `cloned` mirrors
`original` node for node, the base case only has to check one tree running out of nodes;
wherever `original` ends, `cloned` ends at exactly the same moment.

Underneath, it's just a DFS pre-order traversal. The only twist is that it's a paired
traversal over two trees at once, and that's the whole trick: no hash maps, no
bookkeeping to line the two trees up, just two pointers moving together. I like this one
because the part of the problem that looks hard, matching a node's identity across two
separate trees, turns out not to be a real problem once you notice the copy guarantee.

### The solution

![Example 1](images/1.png "Example1")

```go
// getTargetCopy walks original and clones in lockstep: every recursive call
// advances both pointers together, so wherever the traversal is standing in
// original, clones is standing at the structurally-matching node in the
// cloned tree. That's guaranteed because clones is an exact copy of
// original, which lets the search find target's counterpart purely by
// position/value instead of needing to compare node identities.
func getTargetCopy(original, clones, target *TreeNode) *TreeNode {
	// clones mirrors original node-for-node, so clones == nil alone is
	// enough to detect "ran out of tree" along this path; original doesn't
	// need its own nil check.
	if clones == nil {
		return nil
	}
	// Values are guaranteed unique, so a value match means clones is
	// standing exactly where target stands in original — return this
	// cloned-tree node as the answer.
	if clones.Val == target.Val {
		return clones
	}
	// Search the left subtrees of both trees together first.
	n := getTargetCopy(original.Left, clones.Left, target)
	if n != nil {
		return n
	}
	// Not found on the left; whatever the right subtree search finds (or
	// doesn't) is this call's result.
	return getTargetCopy(original.Right, clones.Right, target)
}
```

Walking it through `original = [7,4,3,null,null,6,19]`, `target = 3` (expected: the `3`
node from `cloned`):

- `getTargetCopy(7, 7, target)`: `clones.Val(7) != target.Val(3)` → recurse left first,
  into `(4, 4)`.

  ![Step 1: at the root pair (7,7), values don't match, recurse left into (4,4)](images/walkthrough-1.svg)

- `getTargetCopy(4, 4, target)`: `clones.Val(4) != target.Val(3)`. `4` is a leaf in both
  trees, so recursing left goes to `(nil, nil)` → returns `nil` immediately, and
  recursing right also goes to `(nil, nil)` → returns `nil`. This call returns `nil` —
  the target isn't anywhere under `4`.

  ![Step 2: at pair (4,4), values don't match and both children are nil, so this branch is a dead end — returns nil](images/walkthrough-2.svg)

- Back in `getTargetCopy(7, 7, target)`: the left recursion returned `nil`, so it falls
  through to search the right subtree instead: `(3, 3)`.

- `getTargetCopy(3, 3, target)`: `clones.Val(3) == target.Val(3)` → match! Returns
  `clones` directly — the `3` node from the *cloned* tree — and that result is
  propagated all the way back up as the final answer. ✓

  ![Step 3: at pair (3,3), values match — clones (the cloned-tree node) is returned as the answer](images/walkthrough-3.svg)

**Complexity:** O(n) time — in the worst case (target is the last node visited, or the
tree is degenerate) every node pair gets visited once. O(h) space for the recursion
stack, where h is the tree's height (O(log n) for a balanced tree, O(n) for a completely
skewed one).

Full code: `easy_problems/1301_1400/find_a_corresponding_node_of_a_binary_tree_in_a_clone_of_that_tree/` in the repo.

#DSA #LeetCode #100DaysOfCode #CodingInterview #BinaryTree #DFS #Golang

---

*Solution and code by Archit Agarwal. Write-up drafted with AI assistance from the code and problem statement.*
