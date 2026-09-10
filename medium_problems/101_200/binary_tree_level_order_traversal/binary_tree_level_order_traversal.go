package binary_tree_level_order_traversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func LevelOrder(root *TreeNode) [][]int {
	arr := make([][]int, 0)
	return traversal(root, arr, 0)
}

// traversal is a depth-first walk that still produces level-ordered
// output, because the accumulator is keyed on depth rather than on the
// order nodes are reached. A node appends into the slot for its own
// level, so which branch it arrived from stops mattering.
func traversal(head *TreeNode, arr [][]int, level int) [][]int {
	if head == nil {
		return arr
	}
	// len(arr)-1 is the deepest level that has a slot so far, so this
	// fires once per level: the first node to reach a depth opens it.
	if (len(arr) - 1) < level {
		a := make([]int, 0)
		arr = append(arr, a)
	}
	arr[level] = append(arr[level], head.Val)
	// Left before right is what makes each level read left to right.
	//
	// Both calls take back the returned slice rather than relying on
	// mutation: growing arr above can reallocate its backing array, so a
	// caller that ignored the return value would lose later levels.
	arr = traversal(head.Left, arr, level+1)
	arr = traversal(head.Right, arr, level+1)
	return arr
}
