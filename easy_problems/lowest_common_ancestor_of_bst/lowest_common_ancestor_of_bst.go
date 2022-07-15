package lowest_common_ancestor_of_bst

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func LowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	return check(root, p.Val, q.Val)
}

func check(head *TreeNode, p, q int) *TreeNode {

	if head == nil {
		return nil
	}
	max := maxValue(p, q)
	min := minValue(p, q)

	if head.Val == p || head.Val == q || (head.Val > min && head.Val < max) {
		return head
	}

	if head.Val > max {
		return check(head.Left, p, q)
	}
	return check(head.Right, p, q)

}

func maxValue(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func minValue(a, b int) int {
	if a < b {
		return a
	}
	return b
}
