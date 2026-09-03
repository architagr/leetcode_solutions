# 365 Days of LeetCode Challenge — Day 5/365

## Balanced Binary Tree (Easy)

🔗 [leetcode.com/problems/balanced-binary-tree](https://leetcode.com/problems/balanced-binary-tree/)

Checking height separately at every node feels reasonable until you realize it quietly
costs O(n²) on a skewed tree, since it keeps re-walking the same subtrees. Computing
height and checking balance are really the same bottom-up walk, so doing them together
in one pass drops it to O(n). The part I like: the moment one branch turns out broken,
the recursion can stop and skip the rest of the tree entirely.

Full breakdown in today's newsletter article.

#DSA #LeetCode #100DaysOfCode #Programming #Algorithms #BinaryTree #Recursion
