## Intuition

Two trees are the same when their roots match and their left subtrees are the same and
their right subtrees are the same. That sentence is already the recursion. The only work is
deciding what "roots match" means when one or both of them is nil.

There are four cases for a pair of nodes, and the order the code checks them in is the
part worth getting right:

1. Both nil. Two empty trees are the same. Return true.
2. Exactly one nil. One tree has a node where the other has nothing, so the shapes differ.
   Return false.
3. Both present, different values. Return false.
4. Both present, same value. The answer is whatever the two child comparisons say.

Case 1 has to come before case 2, because "exactly one is nil" is written as
`p == nil || q == nil`, which is also true when both are. Checking both-nil first means
the `||` only ever sees the lopsided case.

Example 2 is the one that trips up a values-only check. `[1,2]` and `[1,null,2]` hold the
same values, and a traversal that just lists values would call them equal. Walking the two
trees in lockstep catches it, because the `2` on one side lines up with a nil on the
other.

This check turned up inside [Day 70: Subtree of Another Tree](https://github.com/architagr/leetcode_solutions/blob/main/easy_problems/501_600/subtree_of_another_tree/)
as a helper. Here it's the whole problem, and it's about to come back as half of Symmetric
Tree.

One thing about how this code is written: it computes `left` and `right` into variables
and returns `left && right`. That means the right subtrees are compared even when the left
ones have already come back false. Writing `return isSameTree(p.Left, q.Left) &&
isSameTree(p.Right, q.Right)` would let Go's `&&` skip the second call. The answer is the
same either way; the short-circuit version just stops sooner on trees that differ early.

**Complexity:**
- Time: O(n), where n is the number of positions the two trees share. Each pair is
  compared once, and as written no pair is skipped.
- Space: O(h) for the recursion stack.
