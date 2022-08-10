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
	mid := len(nums) / 2

	root := new(TreeNode)
	root.Val = nums[mid]
	root.Left = SortedArrayToBST(nums[:mid])
	if mid < len(nums) {
		root.Right = SortedArrayToBST(nums[mid+1:])
	}
	return root
}
