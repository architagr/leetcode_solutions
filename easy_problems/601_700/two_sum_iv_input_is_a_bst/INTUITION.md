## Intuition

Strip away the "BST" label for a moment and this is just the classic **Two Sum**
problem: find two elements that add up to `k`. The one-pass array solution for
Two Sum keeps a hash set of values seen so far and, for each new value `v`, checks
whether `k - v` is already in the set. If it is, the two matching numbers have been
found; otherwise `v` (or its complement) gets recorded for a future match.

This solution ports that exact idea onto a tree by walking it with plain recursion
(preorder: visit the node, then its left subtree, then its right subtree) instead of
iterating an array. The one twist is *what* gets stored: instead of recording each
visited value directly, `find` stores that value's **complement**, `k - root.Val`, in
`hashMap`. Then, for every new node, it checks whether the node's *own* value is
already sitting in `hashMap` as someone else's complement. If it is, some earlier
node `A` previously computed `k - A.Val == root.Val`, which rearranges to
`A.Val + root.Val == k` — exactly the pair being searched for.

Because `hashMap` is a single map passed down through every recursive call (maps are
reference types in Go, so every call shares the same underlying storage), a match can
be found across *any* two nodes in the tree, not just siblings or ancestors — the
complement recorded while visiting one branch is still visible when a completely
different branch is visited later.

The interesting trade-off: this approach completely ignores the fact that the tree is
a **binary search tree**. A plain hash-set two-sum check works on any binary tree,
sorted or not — the BST ordering isn't used anywhere in `find`. That's a valid choice
(it's simpler to reason about than an in-order-traversal two-pointer approach that
*does* exploit the ordering), but it trades away the ability to do the search with
O(h) extra space instead of O(n).

**Complexity:**
- Time: O(n) — every node is visited once, and each map lookup/insert is O(1) on
  average.
- Space: O(n) — `hashMap` can grow to hold one entry per node in the worst case
  (when no pair sums to `k`), plus O(h) for the recursion stack, where `h` is the
  tree's height.
