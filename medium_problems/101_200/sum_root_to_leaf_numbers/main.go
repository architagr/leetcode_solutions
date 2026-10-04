package sumroottoleafnumbers

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var res int

func sumNumbers(root *TreeNode) int {
	res = 0
	process(root, 0)
	return res
}

func process(node *TreeNode, sum int) {
	if node == nil {
		return
	}
	sum = sum*10 + node.Val
	if node.Left == nil && node.Right == nil {
		res += sum
		return
	}
	process(node.Left, sum)
	process(node.Right, sum)
}
