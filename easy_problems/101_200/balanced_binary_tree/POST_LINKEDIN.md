# 365 Days of LeetCode Challenge — Day 5/365

## Balanced Binary Tree (Easy)

🔗 [leetcode.com/problems/balanced-binary-tree](https://leetcode.com/problems/balanced-binary-tree/)

The naive fix — checking height at every node separately — quietly costs you O(n²) on a
skewed tree. Today's problem is a clean example of why: computing height and validating
balance are really the *same* bottom-up walk, and doing them in one pass instead of two
turns it into O(n). Bonus: the moment one branch is broken, the recursion can stop
checking the rest of the tree entirely.

Full breakdown in today's newsletter article ⬇
