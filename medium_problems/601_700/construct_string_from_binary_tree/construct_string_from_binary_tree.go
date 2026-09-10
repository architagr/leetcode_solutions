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

// parse builds the pre-order string for the subtree at node.
//
// The empty string returned for a nil node is not just a stop signal: a
// caller wraps it in parentheses, which is what produces the "()" the
// formatting rules require when a node has a right child and no left
// one. The rule and the base case are the same mechanism.
func parse(node *TreeNode) string {
	if node == nil {
		return ""
	}
	s := strconv.Itoa(node.Val)
	// A leaf skips this block entirely, so it never emits "()".
	if node.Right != nil || node.Left != nil {
		// Unconditional: no nil check on Left. A missing left child makes
		// parse return "", so this same line emits the "()" placeholder.
		// That is why nothing here tests for "right child but no left".
		s += "(" + parse(node.Left) + ")"
		// The mirror case needs no placeholder - "(left)" with nothing
		// after it is already unambiguous - so this one is conditional.
		if node.Right != nil {
			s += "(" + parse(node.Right) + ")"
		}
	}
	return s
}
