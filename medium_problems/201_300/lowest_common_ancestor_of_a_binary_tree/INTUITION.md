## Intuition

This is yesterday's problem with one guarantee removed, and it's worth seeing how much
that guarantee was doing.

## Builds on

- [Day 27: Lowest Common Ancestor of a Binary Search Tree](https://github.com/architagr/leetcode_solutions/blob/main/medium_problems/201_300/lowest_common_ancestor_of_a_binary_search_tree/) — the same question on a BST, where comparing values told you which way to walk. Take the ordering away and none of that survives

In a BST, comparing a target against a node told you which subtree it was in, so the walk
went down one path and never looked at the rest of the tree. Here there is no ordering, so
a comparison tells you nothing about direction. The only way to know whether a target is
under a node is to go and look.

Which flips the whole shape. Yesterday's solution walked downward, deciding as it went.
This one has to walk everything and report upward: each subtree answers "how many of the
two targets are in me?", and the parent adds the reports together.

That count is the entire idea. A node whose subtrees between them contain both targets —
counting itself as part of its own subtree — is a common ancestor. The *lowest* such node
is the first one to reach a count of two on the way up, which is exactly what a post-order
traversal finds first.

Two details in this implementation are worth calling out because they're easy to get
wrong.

The early returns matter. After recursing left, if `ln != nil` the answer has already been
found deeper and is returned immediately without touching the right subtree. Without those,
an ancestor higher up would also see a count of two and overwrite the answer with itself —
producing a common ancestor, just not the lowest one.

The node's own match is folded into `lcount` rather than a separate variable. A node that
is itself a target counts toward its own total, which is what makes "a node can be a
descendant of itself" work: if `p` is an ancestor of `q`, then at `p` the count reaches two
(one for being `p`, one from the subtree holding `q`), and `p` is correctly the answer.

Comparison is on `Val`, as it was yesterday. That relies on values being unique — the
problem guarantees it. With duplicate values the count could be reached by the wrong nodes
entirely, and comparing pointers would be the fix.

**Complexity:**
- Time: O(n) in the worst case — every node may be visited. The early returns prune real
  work once the answer is found, but they don't change the bound.
- Space: O(h) for the recursion stack, where h is the tree's height.
