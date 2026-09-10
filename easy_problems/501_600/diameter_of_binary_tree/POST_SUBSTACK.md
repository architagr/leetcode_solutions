---
meta_title: "Diameter of a binary tree: try every node as the centre"
meta_description: "The longest path need not pass through the root, but through any given node it is just height(left) + height(right). So try every node and keep the best."
tags: [golang, binary-tree, recursion, dfs, leetcode]
---

# Diameter of Binary Tree

*365 Days of LeetCode Challenge — Day 21/365*

🔗 [LeetCode #543](https://leetcode.com/problems/diameter-of-binary-tree/) · Difficulty: Easy

The trap here is the word "diameter." It suggests something centred, and the longest path in a
binary tree very often isn't — it can sit entirely inside one subtree, nowhere near the root.

Once you stop looking for *the* centre and start treating every node as a candidate centre, the
problem stops being about finding a path at all.

Given the `root` of a binary tree, return the length of the **diameter** of the
tree — the longest path between any two nodes, measured in edges. That path may
or may not pass through the root.

!["Example 1"](diamtree.jpg "Example 1")

```
Input: root = [1,2,3,4,5]
Output: 3
Explanation: 3 is the length of the path [4,2,1,3] or [5,2,1,3].
```

---

The reusable move is computing a candidate answer at every node during a traversal you were doing
anyway. The heights are needed regardless; the diameter falls out of them for free.

Full code and the step-by-step walkthrough:
[diameter_of_binary_tree](https://github.com/architagr/leetcode_solutions/blob/main/easy_problems/501_600/diameter_of_binary_tree/SOLUTION.md)

---

*Solution and code by Archit Agarwal. Write-up drafted with AI assistance from the code and problem statement.*
