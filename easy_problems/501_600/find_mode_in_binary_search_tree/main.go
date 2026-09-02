package findmodeinbinarysarchtree

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findMode(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	// Tally how many times every value occurs. This ignores the BST
	// ordering entirely and treats the tree like any binary tree.
	data := make(map[int]int)
	data = getCnt(root, data)

	res := make([]int, 0)
	cnt := 0
	// First pass: find the highest frequency present in the tree.
	for _, v := range data {
		if v > cnt {
			cnt = v
		}
	}

	// Second pass: collect every value whose frequency matches the max,
	// so ties (multiple modes) are all returned, not just one.
	for k, v := range data {
		if v == cnt {
			res = append(res, k)
		}
	}
	return res
}

// getCnt walks the whole tree and increments data[node.Val] for every
// node visited. Since data is a map (reference type), every recursive
// call mutates the same underlying map, so the return value only
// matters to the top-level caller.
func getCnt(root *TreeNode, data map[int]int) map[int]int {
	if root == nil {
		return nil
	}
	data[root.Val]++
	getCnt(root.Left, data)
	getCnt(root.Right, data)
	return data

}
