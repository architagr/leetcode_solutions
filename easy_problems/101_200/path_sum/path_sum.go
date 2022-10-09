package path_sum

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func hasPathSum(root *TreeNode, targetSum int) bool {
	return sum(root, targetSum, 0)
}

func sum(root *TreeNode, target, current int) bool {
	if root == nil {
		return false
	} else if root.Left != nil && root.Right != nil {
		return sum(root.Left, target, current+root.Val) || sum(root.Right, target, current+root.Val)
	} else if root.Left != nil {
		return sum(root.Left, target, current+root.Val)
	} else if root.Right != nil {
		return sum(root.Right, target, current+root.Val)
	}
	return target == root.Val+current
}
