package deletenodeinabst

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}

	// delete from the right subtree
	if key > root.Val {
		root.Right = deleteNode(root.Right, key)
	} else if key < root.Val { // delete from the left subtree
		root.Left = deleteNode(root.Left, key)
	} else { // delete the current node
		// the node is a leaf
		if root.Left == nil && root.Right == nil {
			root = nil
		} else if root.Right != nil { // the node is not a leaf and has a right child
			root.Val = successor(root)
			root.Right = deleteNode(root.Right, root.Val)
		} else { // the node is not a leaf, has no right child, and has a left child
			root.Val = predecessor(root)
			root.Left = deleteNode(root.Left, root.Val)
		}
	}
	return root
}

func successor(root *TreeNode) int {
	root = root.Right
	for root.Left != nil {
		root = root.Left
	}
	return root.Val
}

func predecessor(root *TreeNode) int {
	root = root.Left
	for root.Right != nil {
		root = root.Right
	}
	return root.Val
}
