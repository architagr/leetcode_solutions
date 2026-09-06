# Intuition

For each node, compare its value against the average of its own subtree. The average needs two numbers: the sum of the subtree and how many nodes are in it.

Computing those independently for every node would mean walking each subtree once per node, which is O(n^2) on a skewed tree. But a subtree's sum and size are built out of its children's sums and sizes, so one pass is enough if the information travels the right direction.

## Answers move upward

A node can't work anything out on the way down. It doesn't know what's underneath it yet. On the way back up, though, both children have already reported, and combining their reports is arithmetic:

```
sum   = leftSum + node.Val + rightSum
count = leftCount + rightCount + 1
```

That's post-order traversal: children first, parent second. The recursive call returns a pair rather than a single number, which is the only structural difference from a plain sum-the-tree function.

Once a node has its own sum and count, checking itself is one comparison. Do that check inside the same call and the count of matches accumulates for free, with no second traversal.

## The empty case does real work

`nil` returns `(0, 0)`. It reads like boilerplate but it's what makes the combine step safe: a node with one child gets `(0, 0)` from the missing side, so the formula above works without any special-casing for one-child nodes.

That matters here more than usual because this is a plain binary tree, not a BST, and one-child nodes are common. In the LeetCode example, node 5 has only a right child, and it's one of the nodes that counts.

## Floor division is load-bearing

The problem says the average is rounded down. Go's integer `/` already truncates toward zero, and since values are non-negative here, truncation and floor are the same thing. So `currentNodeSum/countNodes` is the rounded-down average without any extra work.

This isn't a detail you can ignore. Node 5 in the example has sum 11 over 2 nodes. The true average is 5.5, which is not 5, but the floored average is 5, which is. Reach for floats and that node stops counting and the answer comes out wrong.

The one thing to keep an eye on is that `countNodes` is never zero when the division runs, since the `nil` check returns before it. A node always counts itself.

## Complexity

Time is O(n). One visit per node, constant work each.

Space is O(h) for the recursion stack. O(log n) if the tree is balanced, O(n) if it's a chain. Nothing else is allocated; the running match count is a single integer captured by the closure.
