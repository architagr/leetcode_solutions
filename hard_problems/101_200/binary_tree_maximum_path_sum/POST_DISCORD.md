**365 Days of LeetCode Challenge — Day 25/365**
**Binary Tree Maximum Path Sum** (Hard)
🔗 https://leetcode.com/problems/binary-tree-maximum-path-sum/

First hard of the challenge, and the traversal isn't what makes it hard. Each node has to compute two different things, and confusing them is the whole trap.

A path can't fork — a node with three neighbours in it isn't a path. So at the topmost node of a path, both children may contribute. Anywhere else along it, only one can.

That gives every node two roles:
- top of a path: `left + right + val`, a candidate for the answer
- link in some ancestor's path: `max(left, right) + val`, reported upward

Only one can be the return value. The other is recorded through a pointer as a side effect. Return the two-sided value instead and the parent builds a path that forks.

```go
left := recurssivePathSum(root.Left, max)
right := recurssivePathSum(root.Right, max)
// record left+right+val into *max, then:
return maxValue(maxValue(left, right)+root.Val, root.Val)
```

The nested maxValue calls come out to `val + max(0,left) + max(0,right)` for the candidate and `val + max(0,left,right)` for the return — the comparisons against `val` alone are how a negative branch gets dropped.

O(n) time, O(h) space.

Full walkthrough with step-by-step diagrams: https://github.com/architagr/leetcode_solutions/blob/main/hard_problems/101_200/binary_tree_maximum_path_sum/SOLUTION.md
