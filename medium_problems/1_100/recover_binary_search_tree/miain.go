package recoverbinarysearchtree

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func recoverTree(root *TreeNode) {
	inOrderArray := inOrderTravasal(root)

	var first *TreeNode
	var second *TreeNode
	for i := 1; i < len(inOrderArray); i++ {
		if inOrderArray[i-1].Val > inOrderArray[i].Val {
			if first == nil {
				first = inOrderArray[i-1]
				second = inOrderArray[i]
			} else {
				second = inOrderArray[i]
				break
			}
		}
	}
	second.Val, first.Val = first.Val, second.Val
}

func inOrderTravasal(root *TreeNode) []*TreeNode {
	if root == nil {
		return []*TreeNode{}
	}
	res := inOrderTravasal(root.Left)
	res = append(res, root)
	return append(res, inOrderTravasal(root.Right)...)
}
