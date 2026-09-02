**365 Days of LeetCode Challenge — Day 9/365**

**Binary Tree Postorder Traversal** (Easy)
🔗 https://leetcode.com/problems/binary-tree-postorder-traversal/

Postorder just means "children before parent": recurse left, recurse right,
*then* append the current node. The recursion itself guarantees every
descendant is already recorded before a node ever appends its own value — no
extra bookkeeping needed.

Full breakdown in today's newsletter article ⬇️
