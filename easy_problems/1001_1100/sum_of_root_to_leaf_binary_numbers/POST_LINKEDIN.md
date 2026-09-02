**365 Days of LeetCode Challenge — Day 26/365**

**Sum of Root To Leaf Binary Numbers** (Easy)
🔗 https://leetcode.com/problems/sum-of-root-to-leaf-binary-numbers/

Every root-to-leaf path spells out a binary number top to bottom — so instead of
collecting bits and converting at the end, carry a running value down the recursion:
shift it left and drop in each node's bit as you descend. By the time you hit a leaf,
the running value already *is* the answer for that path.

Full breakdown in today's newsletter article ⬇️
