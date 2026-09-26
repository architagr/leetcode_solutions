## Intuition

A node is good when nothing above it on its root path is bigger than it. Checking that
naively means walking the whole path again at every node. But the only fact about the path
the check needs is its largest value, and that can travel down with the recursion: each
call receives the biggest value seen so far on the way to it, compares, and passes a
possibly bigger one on to its children.

So the state flows top-down, the same way the remaining sum did in Path Sum. What flows
back up is just a count.

## Builds on

- [Day 11: Path Sum](https://github.com/architagr/leetcode_solutions/blob/main/easy_problems/101_200/path_sum/) — carrying a running value down the root-to-node path as a function argument, so each node sees the state of its own path and nothing else

The part I think is worth being precise about is why one variable is enough when the tree
branches. `currentMax` is a parameter, so every call gets its own copy. When the call at
`4` raises its max to 4, that affects the calls below `4` and nothing else; the left
branch under `1` was handed 3 and keeps 3. Each branch's max is exactly the max of its own
path, with no undo step needed. An explicit stack would have to store the max alongside
each node to get the same effect.

Two small details. The comparison is `>=`, not `>`, because the rule is "no node greater
than X", so a tie is still good; in Example 1 the bottom-left `3` counts for that reason.
And `goodNodes` seeds the recursion with `root.Val` as the starting max, which makes the
root good automatically. Seeding with the smallest possible int would work too. Seeding
with `root.Val` relies on the constraint that the tree has at least one node, since a nil
root would panic on `root.Val`.

**Complexity:**
- Time: O(n), each node visited once.
- Space: O(h) for the recursion stack.
