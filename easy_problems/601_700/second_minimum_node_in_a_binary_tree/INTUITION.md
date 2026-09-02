## Intuition

At first glance this looks like "find the second smallest value in a binary tree," which
usually screams *collect everything, sort it, take the second element*. But the problem
statement hands us a very unusual constraint that makes a full traversal-and-sort approach
wasteful: **every node with two children satisfies `root.val = min(root.left.val,
root.right.val)`.**

That single rule tells us two important things before we write a line of code:

1. **The root is always the global minimum.** Since every parent's value is the smaller
   of its children's values, the minimum can never live deeper than the root — it
   propagates all the way up. So `root.Val` *is* the first minimum, always. No search
   needed.
2. **A subtree's minimum equals its own root's value.** By the same rule applied
   recursively, whatever value sits at the root of *any* subtree is also the smallest
   value anywhere inside that subtree. Values never decrease as you go down — they only
   stay the same or grow.

That second fact is the key to an efficient solution. If we're walking the tree looking
for the second-smallest value and we land on a node whose value is *already* strictly
greater than the known minimum, we know two things immediately:
- this node's value is a **candidate** for the second minimum, and
- **nothing underneath it can beat that candidate**, because the whole subtree is bounded
  below by this node's own value.

So there's no reason to keep descending into that subtree. We can prune it entirely.

The only reason to keep recursing into a subtree is when the current node's value is
*still equal* to the global minimum — because that means the "real" second-minimum value,
if it exists, must be hiding somewhere further down where the tree first breaks away from
the minimum value.

### Putting it together

- Track the global minimum (`min`, taken directly from `root.Val`).
- Track the best candidate found so far for the second minimum (`ans`, starting at
  infinity).
- Do a DFS: at each node,
  - if the node's value is strictly between `min` and the current best `ans`, it becomes
    the new `ans` — and we stop descending into that branch (pruning, as explained
    above).
  - if the node's value equals `min`, we haven't found where the tree "branches away"
    yet, so keep recursing into both children.
  - otherwise (value ≥ `ans` already), there's nothing better down there either — the
    branch is a dead end for improving `ans`.
- At the end, if `ans` was ever updated, that's the answer; otherwise every value in the
  tree equals `min` and there is no second minimum, so return `-1`.

### Why this works instead of just collecting-and-sorting

A brute-force "collect all values into a set, sort, take the second element" solution
also works and is easy to reason about, but it visits every node unconditionally and pays
for a sort. This approach still visits nodes in the worst case (e.g. every value in the
tree equals `min` except one leaf) — so its worst-case time complexity is the same — but
it usually prunes large parts of the tree early, avoiding both the extra memory of storing
every value and the cost of sorting.

### Complexity

- **Time:** O(n) worst case, where n is the number of nodes — the pruning helps in
  practice but doesn't change the asymptotic bound, since a single "all values equal
  except one deep leaf" tree still forces a full visit down that one path... and in
  general, once we branch away from `min`, the amount of extra work is bounded by that
  visited subtree, so overall each node is visited at most once.
- **Space:** O(h) for the recursion stack, where h is the height of the tree (O(log n)
  for a balanced tree, O(n) worst case for a completely skewed one). No auxiliary
  collection of values is stored.
