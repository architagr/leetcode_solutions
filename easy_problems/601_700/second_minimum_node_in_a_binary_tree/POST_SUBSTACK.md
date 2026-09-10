---
meta_title: "The root is already the minimum, so stop searching for it"
meta_description: "Every parent holds the smaller of its children, so the minimum propagates to the root. Half the problem is solved before the first line of code."
tags: [golang, binary-tree, recursion, dfs, leetcode]
---

# Second Minimum Node In a Binary Tree

*365 Days of LeetCode Challenge — Day 48/365*

🔗 [LeetCode #671](https://leetcode.com/problems/second-minimum-node-in-a-binary-tree/) · Difficulty: Easy

"Second smallest" usually means collect everything, sort, take index one. This problem hands you
a constraint that makes most of that unnecessary, and it's easy to read past.

Every node with children holds the smaller of the two. Follow that upward and one of the two
values you're looking for is already sitting in your hand.

### The problem

You're given a non-empty binary tree with an odd shape rule baked in: every node has
either **zero** or **two** children, and whenever a node has two children, its own value
is the smaller of its two children's values:

```
root.val = min(root.left.val, root.right.val)
```

Given a tree like this, find the **second smallest distinct value** among all the node
values in the tree. If no such value exists, return `-1`.

Two examples from the problem statement:

!["Example 1"](smbt1.jpg "Example 1")

```
Input: root = [2,2,5,null,null,5,7]
Output: 5
Explanation: The smallest value is 2, the second smallest value is 5.
```

!["Example 2"](smbt2.jpg "Example 2")

```
Input: root = [2,2,2]
Output: -1
Explanation: The smallest value is 2, but there isn't any second smallest value.
```

---

The habit worth building is reading unusual constraints as *given answers* rather than as
background. This one collapses "find the two smallest" into "find the smallest value strictly
greater than the root," which is a different and much cheaper search.

Full code and the step-by-step walkthrough:
[second_minimum_node_in_a_binary_tree](https://github.com/architagr/leetcode_solutions/blob/main/easy_problems/601_700/second_minimum_node_in_a_binary_tree/SOLUTION.md)

---

*Solution and code by Archit Agarwal. Write-up drafted with AI assistance from the code and problem statement.*
