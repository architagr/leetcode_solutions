**365 Days of LeetCode Challenge — Day 2/365**
**Convert Sorted Array to Binary Search Tree** (Easy)

Sorted array → balanced BST: pick the middle element as root, recurse left/right on
the two halves. Balance falls out for free since the array's already sorted.

```go
mid := len(nums) / 2
root := &TreeNode{Val: nums[mid]}
root.Left = build(nums[:mid])
root.Right = build(nums[mid+1:])
```

🔗 https://leetcode.com/problems/convert-sorted-array-to-binary-search-tree/
