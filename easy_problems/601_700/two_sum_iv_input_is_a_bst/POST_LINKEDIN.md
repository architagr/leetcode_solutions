**365 Days of LeetCode Challenge — Day 17/365**

**Two Sum IV - Input is a BST** (Easy)
🔗 https://leetcode.com/problems/two-sum-iv-input-is-a-bst/

Strip away the "BST" and this is plain Two Sum: keep a hash map of complements
as you walk the tree, and check each new node against it. The fun part is the
shared map means a match can turn up between two completely unrelated branches
— and the code still checks the right subtree even after the left one already
found the answer.

Full breakdown in today's newsletter article ⬇️
