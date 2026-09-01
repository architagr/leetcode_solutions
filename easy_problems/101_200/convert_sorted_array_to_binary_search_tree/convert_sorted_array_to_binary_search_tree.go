package convert_sorted_array_to_binary_search_tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

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
	// mid is always < len(nums) here since nums is non-empty, so this
	// recurses on the right half; nums[mid+1:] is already an empty slice
	// when there's nothing left, so the base case handles it naturally.
	if mid < len(nums) {
		root.Right = SortedArrayToBST(nums[mid+1:])
	}
	return root
}
