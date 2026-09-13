package allelementsintwobinarysearchtrees

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func getAllElements(root1, root2 *TreeNode) []int {
	root1Inorder := inorder(root1, []int{})
	root2Inorder := inorder(root2, []int{})
	i, j := 0, 0
	result := make([]int, 0, len(root1Inorder)+len(root2Inorder))

	for i < len(root1Inorder) && j < len(root2Inorder) {
		if root1Inorder[i] < root2Inorder[j] {
			result = append(result, root1Inorder[i])
			i++
		} else {
			result = append(result, root2Inorder[j])
			j++
		}
	}
	for ; i < len(root1Inorder); i++ {
		result = append(result, root1Inorder[i])
	}
	for ; j < len(root2Inorder); j++ {
		result = append(result, root2Inorder[j])
	}
	return result
}

func inorder(root *TreeNode, arr []int) []int {
	if root == nil {
		return arr
	}
	arr = inorder(root.Left, arr)
	arr = append(arr, root.Val)
	arr = inorder(root.Right, arr)
	return arr
}
