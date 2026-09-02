package average_of_levels_in_binary_tree

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func AverageOfLevel(root *TreeNode) []float64 {
	// result[i] will hold the sum of node values at level i, count[i] the
	// number of nodes at level i. Passed by pointer so every recursive call
	// shares (and can grow) the same underlying slices.
	result := make([]float64, 0)
	count := make([]float64, 0)
	getSumAndCount(root, &result, &count, 0)
	// Convert each level's running sum into that level's average in place.
	for i := range result {
		result[i] /= count[i]
	}
	return result
}
func getSumAndCount(node *TreeNode, result, count *[]float64, level int) {
	if node == nil {
		return
	}
	// First time we reach this depth: lazily grow both slices by one slot.
	// Later visits to the same depth (from other branches) reuse this slot.
	if len(*result) < level+1 {
		*result = append(*result, 0.0)
		*count = append(*count, 0.0)
	}

	// Indexing by level (not visit order) is what lets nodes from different
	// branches of the tree accumulate into the same slot as long as they
	// share a depth.
	(*result)[level] += float64(node.Val)
	(*count)[level]++
	// Preorder DFS: left subtree fully explored before the right one.
	getSumAndCount(node.Left, result, count, level+1)
	getSumAndCount(node.Right, result, count, level+1)
}
