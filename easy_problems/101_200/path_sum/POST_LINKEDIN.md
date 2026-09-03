**365 Days of LeetCode Challenge — Day 7/365**

**Path Sum** (Easy)
🔗 https://leetcode.com/problems/path-sum/

Root-to-leaf is the whole catch on this one. A node with children never gets checked
against the target, even if the running sum happens to land on it along the way, only
a genuine leaf does. Carry the sum down the recursion as you descend, and compare it
exactly once, right when you run out of children to hand it to.

Full walkthrough, the code, and the recursion trace are in today's newsletter
article, linked below.

#DSA #LeetCode #100DaysOfCode #SoftwareEngineering #BinaryTree #DFS #Golang
