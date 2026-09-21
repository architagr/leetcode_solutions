package constructbinarytreefrompreorderandpostordertraversal

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func constructFromPrePost(preorder []int, postorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	root := &TreeNode{Val: preorder[0]}
	if len(preorder) == 1 {
		return root
	}
	i := 0
	for ; i < len(postorder); i++ {
		if postorder[i] == preorder[1] {
			break
		}
	}
	root.Left = constructFromPrePost(preorder[1:i+2], postorder[:i+1])
	root.Right = constructFromPrePost(preorder[i+2:], postorder[i+1:len(postorder)-1])
	return root
}
