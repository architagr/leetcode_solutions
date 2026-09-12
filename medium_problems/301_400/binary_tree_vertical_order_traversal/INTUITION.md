## Intuition

Give every node a column number. The root is column `0`, a left child is one column to the
left of its parent, a right child one to the right. Group by that number and you have the
answer.

That much is a coordinate carried down the traversal, exactly like the depth parameter from
Day 29 — the only difference is that this one can go negative, because columns extend both
ways from the root.

## Builds on

- [Day 29: Binary Tree Level Order Traversal](https://github.com/architagr/leetcode_solutions/blob/main/medium_problems/101_200/binary_tree_level_order_traversal/) — grouping by a coordinate carried down the traversal, and the BFS-versus-DFS question that this problem answers the other way

The interesting part is that here the traversal *has* to be BFS, and Day 29's did not.

Day 29 grouped by depth, and within a level the required order — left to right — came from
recursing into `Left` before `Right`. Arrival order across branches never mattered, so a
depth-first walk worked.

This problem asks for each column top to bottom, and left to right among nodes tied in the
same row. Those are exactly the two orderings a breadth-first walk produces: it visits by
increasing depth, and within a depth, left to right. Append to a column's list as you pop
nodes and the list is already correct, with no sorting.

Swap in a DFS and it breaks. A depth-first walk would drive one branch to the bottom before
touching the other, so a deep node in the left subtree would be appended to its column
before a shallow node from the right subtree that belongs above it. Same grouping, wrong
order inside each group.

So the two problems land on opposite answers to the same question, and the deciding factor
isn't the grouping — it's whether the order *within* a group depends on when nodes are
reached.

The last step assembles the columns in left-to-right order, which is the one place this
implementation takes a shortcut worth naming: it loops `i` from `-101` to `101` and picks up
whatever keys exist. That's leaning on the constraint that the tree holds at most 100 nodes,
so no column index can be outside that span. Collecting the map's keys and sorting them
would be the version that doesn't depend on the constraint.

**Complexity:**
- Time: O(n) for the traversal. The assembly loop is a fixed 203 iterations here, which is
  O(1) given the constraint; the sorted-keys version would be O(c log c) in the number of
  columns.
- Space: O(n) for the map and the queue.
