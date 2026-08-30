package increaseordersearchtree

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func increasingBST(root *TreeNode) *TreeNode {
	inorder := inOrder(root)
	for i := 1; i < len(inorder); i++ {
		inorder[i-1].Right = inorder[i]
	}
	return inorder[0]
}

func inOrder(root *TreeNode) []*TreeNode {
	if root == nil {
		return []*TreeNode{}
	}
	res := make([]*TreeNode, 0)
	left := inOrder(root.Left)
	right := inOrder(root.Right)

	root.Left = nil
	root.Right = nil
	res = append(left, root)
	res = append(res, right...)
	return res
}
