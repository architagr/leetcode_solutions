package construct_b_tree_preorder_traversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildNode(value int) *TreeNode {
	node := new(TreeNode)
	node.Val = value

	return node
}
func BuildBinaryTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	leftnodeIndex := 0
	leftNodeValue := preorder[0]
	root := buildNode(leftNodeValue)
	leninOrder := len(inorder)
	for i := 0; i < leninOrder; i++ {
		if inorder[i] == leftNodeValue {
			leftnodeIndex = i
			break
		}
	}
	if len(preorder) > 0 {
		root.Left = BuildBinaryTree(preorder[1:leftnodeIndex+1], inorder[:leftnodeIndex])
	}
	if leninOrder > 0 {
		root.Right = BuildBinaryTree(preorder[leftnodeIndex+1:], inorder[leftnodeIndex+1:])
	}
	return root
}
