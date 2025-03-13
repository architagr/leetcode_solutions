package diameterofbinarytree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func diameterOfBinaryTree(root *TreeNode) int {
	dia := 0
	calc(root, &dia)
	return dia
}

func calc(root *TreeNode, dia *int) int {
	if root == nil {
		return 0
	}
	left := calc(root.Left, dia)
	right := calc(root.Right, dia)
	*dia = maxVal(left+right, *dia)
	return maxVal(left, right) + 1
}
func maxVal(a, b int) int {
	if a > b {
		return a
	}
	return b
}
