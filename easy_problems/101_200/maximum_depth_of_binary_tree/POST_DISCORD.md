**365 Days of LeetCode Challenge — Day 1/365**
**Maximum Depth of Binary Tree** (Easy)

BFS with a twist: push a `nil` sentinel after the root to mark end-of-level. Pop a
sentinel → level done, bump depth counter, push the next sentinel if nodes remain.

```go
queue := []*TreeNode{root, nil}
// pop nil -> level complete, max++, push new nil if queue non-empty
```

🔗 https://leetcode.com/problems/maximum-depth-of-binary-tree/
