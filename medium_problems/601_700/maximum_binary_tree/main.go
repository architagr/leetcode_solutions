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
