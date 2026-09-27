package maximumbinarytree

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func constructMaximumBinaryTree(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	// The definition, written as recursion. Each call scans its slice for the
	// maximum, so a sorted input costs O(n^2); a monotonic stack does it in O(n).
	higestVal, higestIndex := findHigest(nums)
	return &TreeNode{
		Val:   higestVal,
		Left:  constructMaximumBinaryTree(nums[:higestIndex]),
		Right: constructMaximumBinaryTree(nums[higestIndex+1:]),
	}
}

func findHigest(nums []int) (higestVal, higestIndex int) {
	higestVal = -1
	for start := 0; start < len(nums); start++ {
		if nums[start] > higestVal {
			higestVal = nums[start]
			higestIndex = start
		}
	}
	return
}
