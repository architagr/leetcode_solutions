## Intuition

A tree is symmetric when its right half is the mirror image of its left half. Mirror the
right half and it should be identical to the left. Both of those operations were the last
two days in the series, so this solution is those two functions called one after the other:

```go
return isSameTree(root.Left, invertTree(root.Right))
```

## Builds on

- [Day 84: Same Tree](https://github.com/architagr/leetcode_solutions/blob/main/easy_problems/1_100/same_tree/) — the lockstep comparison of two trees, nil cases and all
- [Day 85: Invert Binary Tree](https://github.com/architagr/leetcode_solutions/blob/main/easy_problems/201_300/invert_binary_tree/) — mirrors a subtree in place by swapping children at every node

I like this as a reading exercise, because it states the definition almost directly. It
also has a cost that's easy to miss, and I'd want to mention it if I wrote this in an
interview.

`invertTree` works in place. Calling it on `root.Right` rewrites the caller's tree: after
`isSymmetric` returns, the right half is still mirrored. For a function whose name starts
with "is", that's a surprising side effect, and it isn't harmless. Call it twice on
Example 1 and you get `true`, then `false`: the second call inverts the right half back to
its original shape and compares `[2,3,4]` against `[2,4,3]`.

The version that avoids it compares the two halves as mirrors directly, without building
one: two subtrees are mirrors when their roots match, the left of one mirrors the right of
the other, and the right of one mirrors the left of the other. Same O(n), one pass instead
of two, and nothing mutated. The follow-up's iterative version is the same idea with a
queue of node pairs.

Given the constraints (at least one node), the `root == nil` check is defensive, but it's
the right answer for an empty tree anyway.

**Complexity:**
- Time: O(n). `invertTree` visits the right half once and `isSameTree` walks both halves
  once, so about 1.5n node visits.
- Space: O(h) for the recursion stack.
