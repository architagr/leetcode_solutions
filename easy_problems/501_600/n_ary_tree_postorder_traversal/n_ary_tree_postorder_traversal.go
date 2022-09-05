package n_ary_tree_postorder_traversal

type Node struct {
	Val      int
	Children []*Node
}

func Postorder(root *Node) []int {
	result := make([]int, 0)
	parse(&result, root)
	return result
}
func parse(result *[]int, node *Node) {
	if node == nil {
		return
	}

	for i := 0; i < len(node.Children); i++ {
		parse(result, node.Children[i])
	}
	(*result) = append((*result), node.Val)
}
