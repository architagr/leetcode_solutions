## 365 Days of LeetCode Challenge — Day 2/365

# Convert Sorted Array to Binary Search Tree

🔗 https://leetcode.com/problems/convert-sorted-array-to-binary-search-tree/ · Difficulty: Easy

### The problem

Given an integer array `nums` sorted in ascending order, convert it to a
**height-balanced** binary search tree.

### The intuition

A height-balanced BST needs roughly equal numbers of elements on the left and right of
every node. Since the input array is already **sorted**, that balance falls out almost
for free: if you always pick the **middle element** of a sorted slice as a node's value,
everything to its left is smaller (goes in the left subtree) and everything to its
right is larger (goes in the right subtree) — and both halves are roughly the same
size.

So the approach is a straightforward divide-and-conquer:
- Pick the middle element of the current slice as the root of this subtree.
- Recurse on the left half to build the left subtree.
- Recurse on the right half to build the right subtree.
- An empty slice means "no subtree here" (`nil`).

Because the recursion always splits the array roughly in half at each step, the
resulting tree is automatically balanced — no rebalancing needed afterward.

### The solution

LeetCode's own example illustrates the target shape for `nums = [-10,-3,0,5,9]`:

![Example 1](images/1.jpg "Example1")

The trace below builds the alternate accepted shape `[0,-3,9,-10,null,5]`, shown in
LeetCode's own alternative diagram:

![Alternative valid answer](images/2.jpg "Example1 alternative")

```go
func SortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	// The array is sorted, so the middle element naturally splits the
	// remaining elements evenly between the left and right subtrees,
	// keeping the resulting tree height-balanced with no extra rebalancing.
	mid := len(nums) / 2

	root := new(TreeNode)
	root.Val = nums[mid]
	root.Left = SortedArrayToBST(nums[:mid])
	root.Right = SortedArrayToBST(nums[mid+1:])
	return root
}
```

Walking it through `nums = [-10,-3,0,5,9]`:
- `mid = 2`, `root.Val = 0`.

  ![Step 1: mid=2, root=0](images/walkthrough-1.svg)
- Left: `SortedArrayToBST([-10,-3])` → root `-3` with left child `-10`.

  ![Step 2: left half [-10,-3], mid=1, root=-3](images/walkthrough-2.svg)
- Right: `SortedArrayToBST([5,9])` → root `9` with left child `5`.

  ![Step 3: right half [5,9], mid=1, root=9](images/walkthrough-3.svg)
- Final tree: `0` with left subtree `-3`(`-10`) and right subtree `9`(`5`) — matching
  the example's accepted output shape `[0,-3,9,-10,null,5]`.

**Complexity:** O(n) time — each element becomes exactly one node. O(log n) recursion
depth plus O(n) for the output tree.

Full code: `easy_problems/101_200/convert_sorted_array_to_binary_search_tree/` in the repo.
