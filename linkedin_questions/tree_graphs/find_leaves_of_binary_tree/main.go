package findleavesofbinarytree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findLeaves(root *TreeNode) [][]int {
	arr := make([][]int, 0)
	_ = parse(root, &arr)
	return arr
}
func parse(root *TreeNode, arr *[][]int) int {
	if root == nil {
		return -1
	}
	leftIndex := parse(root.Left, arr)
	rightIndex := parse(root.Right, arr)
	index := maxVal(leftIndex, rightIndex) + 1
	if len(*arr) <= index {
		*arr = append(*arr, []int{})
	}
	(*arr)[index] = append((*arr)[index], root.Val)
	return index
}

func maxVal(a, b int) int {
	if a > b {
		return a
	}
	return b
}
