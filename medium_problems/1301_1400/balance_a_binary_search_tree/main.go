package balanceabinarysearchtree

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func balanceBST(root *TreeNode) *TreeNode {
	arr := make([]int, 0, 10_000)
	arr = inOrder(root, arr)
	return createNode(arr)
}
func inOrder(node *TreeNode, arr []int) []int {
	if node == nil {
		return arr
	}

	arr = inOrder(node.Left, arr)
	arr = append(arr, node.Val)
	arr = inOrder(node.Right, arr)
	return arr
}

func createNode(arr []int) *TreeNode {
	if len(arr) == 0 {
		return nil
	}

	l := len(arr)
	mid := l / 2
	return &TreeNode{
		Val:   arr[mid],
		Left:  createNode(arr[:mid]),
		Right: createNode(arr[mid+1:]),
	}
}
