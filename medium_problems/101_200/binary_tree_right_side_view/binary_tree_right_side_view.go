package binary_tree_right_side_view

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func BinaryTreeRightSideView(root *TreeNode) []int {
	// Capacity 100 because the constraints cap the tree at 100 nodes, so
	// the slice never has to grow.
	arr := make([]int, 0, 100)

	if root == nil {
		return arr
	}
	arr = findRight(root, arr, 0)
	return arr
}

func findRight(head *TreeNode, arr []int, level int) []int {

	if head == nil {
		return arr
	}
	// arr holds one value per level filled so far, so len(arr) is the
	// next unrecorded level. This is therefore true only for the FIRST
	// node reached at this depth; later arrivals append nothing.
	if len(arr) == level {
		arr = append(arr, head.Val)
	}

	// Right before left, and this ordering is the whole solution: it is
	// what makes the first arrival at a depth the rightmost node there.
	// Swap these two lines and the function returns the left side view.
	//
	// The left subtree is still walked because the rightmost visible node
	// at a level need not be in the right subtree - the right branch can
	// run out of depth first.
	arr = findRight(head.Right, arr, level+1)
	arr = findRight(head.Left, arr, level+1)
	return arr
}
