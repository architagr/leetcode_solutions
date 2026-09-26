package allelementsintwobinarysearchtrees

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// getAllElements reads each BST in order, which yields two sorted slices,
// and merges them the way merge sort does. Using the BST ordering is what
// avoids a sort of the combined values.
func getAllElements(root1, root2 *TreeNode) []int {
	root1Inorder := inorder(root1, []int{})
	root2Inorder := inorder(root2, []int{})
	i, j := 0, 0
	// The output size is known up front, so the appends below never reallocate.
	result := make([]int, 0, len(root1Inorder)+len(root2Inorder))

	for i < len(root1Inorder) && j < len(root2Inorder) {
		// Strict < sends ties to root2. For bare ints either order is correct.
		if root1Inorder[i] < root2Inorder[j] {
			result = append(result, root1Inorder[i])
			i++
		} else {
			result = append(result, root2Inorder[j])
			j++
		}
	}
	// One side is used up. The other's remainder is sorted and larger than
	// everything taken so far, so it is copied as is. At most one of these
	// two loops runs.
	for ; i < len(root1Inorder); i++ {
		result = append(result, root1Inorder[i])
	}
	for ; j < len(root2Inorder); j++ {
		result = append(result, root2Inorder[j])
	}
	return result
}

// inorder appends the subtree's values in sorted order. The slice is
// returned rather than appended to in place because append may reallocate,
// and the caller has to keep the slice that actually holds the values.
func inorder(root *TreeNode, arr []int) []int {
	if root == nil {
		return arr
	}
	arr = inorder(root.Left, arr)
	arr = append(arr, root.Val)
	arr = inorder(root.Right, arr)
	return arr
}
