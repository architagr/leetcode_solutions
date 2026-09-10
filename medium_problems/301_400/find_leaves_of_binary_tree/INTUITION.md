## Intuition

The problem describes a process: strip the leaves, strip the new leaves, repeat. Simulating
that literally means walking the whole tree once per round, which is O(n * h).

You don't have to. A node's round is decided before you remove anything, and it isn't the
node's depth.

Every per-level problem in this batch so far grouped nodes by depth from the root — level
order, right side view, level sums. This one groups by the opposite measurement: distance
down to the deepest leaf beneath a node. Its height, not its depth.

A leaf has nothing below it, so it goes in round 0. A node goes in the round after both its
children have gone, which means its round is one more than the later of its two children's:

```
round(node) = max(round(left), round(right)) + 1
```

That's a height calculation, and it's computed bottom-up in a single post-order pass. The
recursion returns each node's round to its parent, and on the way past, appends the node's
value into the bucket for that round.

## Builds on

- [Day 1: Maximum Depth of Binary Tree](https://github.com/architagr/leetcode_solutions/blob/main/easy_problems/101_200/maximum_depth_of_binary_tree/) — the height recursion this is, with the height put to a different use
- [Day 21: Diameter of Binary Tree](https://github.com/architagr/leetcode_solutions/blob/main/easy_problems/501_600/diameter_of_binary_tree/) — the same post-order shape, where each call returns a height to its parent and does its own work on the way past

Two details make it work cleanly.

The nil case returns `-1`, not `0`. That looks off by one until you follow it: a leaf's two
children are both nil, so `max(-1, -1) + 1` is `0`, and leaves land in round 0 where they
belong. Return `0` for nil and every node's round shifts up by one, and there'd be an empty
first bucket.

The outer slice grows lazily, the same way Day 15's did. `if len(*arr) <= index` creates the
bucket the first time any node reaches that round.

Worth noticing that the answer's shape falls out for free. Nodes appear in `arr[index]` in
whatever order the traversal reaches them, and the problem explicitly says order within a
round doesn't matter — so no sorting, no second pass.

**Complexity:**
- Time: O(n). Every node is visited once. The literal simulation would be O(n * h), so this
  is the whole reason not to simulate.
- Space: O(h) for the recursion stack plus O(n) for the output. `arr` is passed as a
  pointer so the growth in one call is visible to every other.
