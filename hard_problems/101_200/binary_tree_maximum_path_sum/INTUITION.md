## Intuition

The first hard of the challenge, and what makes it hard isn't the traversal. It's that
each node has to compute two different things, and confusing them is the whole trap.

A path here is a sequence of nodes connected by edges, each used once, and it doesn't have
to pass through the root. Read that carefully and one consequence follows: at the topmost
node of a path, the path may come up one side and go down the other. Anywhere else along
the path, it can only continue in one direction — a path that branched would have a node
with three neighbours in it, which isn't a path.

So every node plays two roles:

- **As the top of a path.** Both children can contribute. That's `left + right + val`, and
  it's a candidate for the answer.
- **As a link in some ancestor's path.** Only one child can contribute, because the
  ancestor will also be attached above. That's `max(left, right) + val`, and it's what the
  node reports upward.

Those two values differ, and only one of them can be the return value. The other has to
go somewhere else, which is why `max` is threaded through as a pointer: the answer is
recorded as a side effect at every node, while the return value carries something
different.

## Builds on

- [Day 13: Diameter of Binary Tree](https://github.com/architagr/leetcode_solutions/blob/main/easy_problems/501_600/diameter_of_binary_tree/) — the identical shape. There, each node returned a height upward while recording left + right as a candidate answer. Swap "height" for "path sum" and you have this problem

The other half is negatives. Node values can be negative, so a subtree isn't automatically
worth attaching. Every use of a child's value here is guarded by a comparison against
taking nothing at all, which is what `maxValue(left+root.Val, root.Val)` is doing — it
means "extend into the left child, or don't."

The nested `maxValue` calls in the code look worse than they are. Written out, the
candidate is:

```
val + max(0, left) + max(0, right)
```

and the return is:

```
val + max(0, left, right)
```

The implementation reaches those same two values by a longer route — comparing `val`
against `left+val`, `right+val` and `left+right+val` — but they agree on every input.

**Complexity:**
- Time: O(n) — every node is visited once and does constant work.
- Space: O(h) for the recursion stack, where h is the tree's height.
