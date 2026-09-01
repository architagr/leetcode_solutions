**365 Days of LeetCode Challenge — Day 2/365**

**Convert Sorted Array to Binary Search Tree** (Easy)
🔗 https://leetcode.com/problems/convert-sorted-array-to-binary-search-tree/

Turn a sorted array into a height-balanced BST with zero rebalancing needed. The trick:
since the array is already sorted, the middle element always has smaller values to its
left and larger to its right — pick it as the root, recurse on both halves, and the
tree comes out balanced for free.

```go
func SortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	mid := len(nums) / 2
	root := &TreeNode{Val: nums[mid]}
	root.Left = SortedArrayToBST(nums[:mid])
	root.Right = SortedArrayToBST(nums[mid+1:])
	return root
}
```

Full solution + walkthrough: `easy_problems/101_200/convert_sorted_array_to_binary_search_tree/`
