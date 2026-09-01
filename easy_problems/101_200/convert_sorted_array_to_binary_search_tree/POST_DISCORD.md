**365 Days of LeetCode Challenge — Day 2/365**
**Convert Sorted Array to Binary Search Tree** (Easy)
🔗 https://leetcode.com/problems/convert-sorted-array-to-binary-search-tree/

**Intuition:** Since the array's sorted, the middle element always has smaller values
to its left and larger to its right. Pick it as root, recurse on both halves — the
tree comes out balanced for free, no rebalancing needed.

**Full solution:**
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

**Walkthrough** on `[-10,-3,0,5,9]`:
- `mid=2`, root `0`
- left: `SortedArrayToBST([-10,-3])` → root `-3`, left child `-10`
- right: `SortedArrayToBST([5,9])` → root `9`, left child `5`
- **final tree:** `0` with left subtree `-3(-10)`, right subtree `9(5)` — matches
  `[0,-3,9,-10,null,5]` ✓

O(n) time, O(log n) recursion depth.
