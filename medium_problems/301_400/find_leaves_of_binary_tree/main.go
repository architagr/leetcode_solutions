package findleavesofbinarytree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findLeaves(root *TreeNode) [][]int {
	arr := make([][]int, 0)
	// parse's return value is a round number, which matters to a parent
	// but not here: the answer accumulates inside arr as it goes.
	_ = parse(root, &arr)
	return arr
}

// parse returns the round in which root is removed, and appends root's
// value into that round's bucket on the way past.
//
// A node's round is its HEIGHT - the distance down to its deepest leaf -
// not its depth from the root. That is why a shallow leaf shares round 0
// with a deep one.
//
// arr is a pointer because growing the outer slice in one call has to be
// visible to every other call.
func parse(root *TreeNode, arr *[][]int) int {
	// -1, not 0: a leaf's two nil children make max(-1,-1)+1 = 0, which
	// puts leaves in round 0. Returning 0 here would shift every round up
	// by one and leave an empty bucket at the front.
	if root == nil {
		return -1
	}
	leftIndex := parse(root.Left, arr)
	rightIndex := parse(root.Right, arr)
	// max, because a node is removed the round after BOTH children have
	// gone, so it waits for the slower side.
	index := maxVal(leftIndex, rightIndex) + 1
	if len(*arr) <= index {
		*arr = append(*arr, []int{})
	}
	(*arr)[index] = append((*arr)[index], root.Val)
	return index
}

func maxVal(a, b int) int {
	if a > b {
		return a
	}
	return b
}
