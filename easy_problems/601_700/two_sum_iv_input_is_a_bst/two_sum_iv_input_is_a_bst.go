package two_sum_iv_input_is_a_bst

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findTarget(root *TreeNode, k int) bool {
	var hashMap map[int]bool = make(map[int]bool)

	return find(root, k, hashMap)
}

func find(root *TreeNode, k int, hashMap map[int]bool) bool {
	if root == nil {
		return false
	}
	if _, ok := hashMap[root.Val]; ok {
		return true
	}
	hashMap[k-root.Val] = true
	left := find(root.Left, k, hashMap)
	right := find(root.Right, k, hashMap)
	return left || right
}
