package n_ary_tree_preorder_traversal

type Node struct {
	Val      int
	Children []*Node
}

func Preorder(root *Node) []int {
	arr := make([]int, 0)
	return traversal(root, arr)
}

func traversal(A *Node, arr []int) []int {
	if A == nil {
		return arr
	}
	arr = append(arr, A.Val)
	for _, node := range A.Children {
		arr = traversal(node, arr)
	}
	return arr
}
