## Intuition

At first glance this looks like the mirror image of Maximum Depth of Binary Tree —
just swap `max` for `min` and you're done, right? That's the trap. Minimum depth
isn't "the shallowest point in the tree," it's **the shortest path down to a leaf**,
and a leaf is specifically defined as a node with *no children at all*.

That distinction matters because a node with only one child is not a leaf, even
though one of its subtrees is empty. If you naively took
`min(minDepth(root.Left), minDepth(root.Right)) + 1` at every node, a node with a
missing left child would report a left-subtree depth of `0`, and `min(0, right) + 1`
would collapse to `1` — claiming the tree bottoms out one step below a node that
isn't actually a leaf. That's wrong: you're required to keep walking down the side
that *does* exist until you hit a real leaf.

So the fix is to treat "child is nil" as "that side doesn't count," not as "that side
has depth zero and therefore wins the min." Concretely:
- If both children exist (both recursive depths come back non-zero), it's a genuine
  branching point — take the smaller of the two subtree depths and add one for the
  current node.
- If only one child exists, that's the only path down, regardless of whether it's
  shorter or longer than a nonexistent sibling — follow it and add one.
- The base case (`root == nil`) returns `0`, which is what lets the parent detect "no
  subtree here" versus a real subtree.

This one adjustment — refusing to let a missing child masquerade as a zero-length
path — is the entire difference between this problem and Maximum Depth.

**Complexity:**
- Time: O(n) — every node is visited exactly once.
- Space: O(h) for the recursion stack, where h is the tree's height (O(log n) for a
  balanced tree, O(n) for a completely skewed one).
