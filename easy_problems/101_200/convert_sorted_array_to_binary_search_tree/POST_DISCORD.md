**365 Days of LeetCode Challenge — Day 2/365**
**Convert Sorted Array to Binary Search Tree** (Easy)
🔗 https://leetcode.com/problems/convert-sorted-array-to-binary-search-tree/

Since the array's sorted, the middle element always has smaller values to its left and larger to its right. Pick it as root, recurse on both halves, and the tree comes out balanced with no extra work. Kind of a nice freebie for an easy-rated problem.

```go
func SortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	mid := len(nums) / 2

	root := new(TreeNode)
	root.Val = nums[mid]
	root.Left = SortedArrayToBST(nums[:mid])
	if mid < len(nums) {
		root.Right = SortedArrayToBST(nums[mid+1:])
	}
	return root
}
```

Full walkthrough with step-by-step diagrams: https://github.com/architagr/leetcode_solutions/blob/main/easy_problems/101_200/convert_sorted_array_to_binary_search_tree/SOLUTION.md
