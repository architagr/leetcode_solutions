package flatten_binary_tree_to_linked_list

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Flatten rewrites root into a right-leaning chain in pre-order, in two
// separate passes: collect the values, then rebuild. The problem's
// follow-up asks for an in-place O(1)-space version; this trades that
// space for code that can be read in one go.
func Flatten(root *TreeNode) {
	if root == nil {
		return
	}
	temp := root
	arr := make([]TreeNode, 0)
	getPreOrderArray(temp, &arr)

	// arr[0] is safe to index without a length check: root is non-nil,
	// so the traversal appended at least one value.
	//
	// The root keeps its identity because the caller holds this pointer,
	// so its value is overwritten rather than the node being replaced.
	root.Val = arr[0].Val
	temp = root
	temp.Left = nil

	// Build the chain front to back, temp always pointing at its tail.
	// The final node's Right is never assigned, and new(TreeNode) zeroed
	// it, so the list terminates without an explicit step.
	for i := 1; i < len(arr); i++ {
		y := new(TreeNode)
		y.Val = arr[i].Val
		temp.Right = y
		// Never load-bearing - the root was cleared above and every
		// other temp came from new(TreeNode) - but it states the loop's
		// invariant: nothing this loop leaves behind has a left child.
		temp.Left = nil
		temp = temp.Right
	}

}

// getPreOrderArray appends every value in pre-order.
//
// It collects TreeNode values carrying only Val, not the nodes
// themselves - the rebuild allocates fresh nodes, so nothing but the
// numbers crosses between the two phases. arr is taken as a pointer so
// an append that reallocates is visible to every caller.
func getPreOrderArray(root *TreeNode, arr *[]TreeNode) {

	if root == nil {
		return
	}

	(*arr) = append((*arr), TreeNode{
		Val: root.Val,
	})
	getPreOrderArray(root.Left, arr)
	getPreOrderArray(root.Right, arr)

}
