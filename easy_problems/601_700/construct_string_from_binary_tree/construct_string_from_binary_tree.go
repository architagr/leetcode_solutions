package construct_string_from_binary_tree

import "strconv"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func Tree2str(root *TreeNode) string {
	return parse(root)
}

func parse(node *TreeNode) string {
	if node == nil {
		return ""
	}
	s := strconv.Itoa(node.Val)
	if node.Right != nil || node.Left != nil {
		s += "(" + parse(node.Left) + ")"
		if node.Right != nil {
			s += "(" + parse(node.Right) + ")"
		}
	}
	return s
}
