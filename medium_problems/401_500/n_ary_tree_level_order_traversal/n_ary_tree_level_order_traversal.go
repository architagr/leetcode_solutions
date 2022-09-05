package n_ary_tree_level_order_traversal

type Node struct {
	Val      int
	Children []*Node
}

func LevelOrder(root *Node) [][]int {
	result := make([][]int, 0)
	parser(&result, 0, root)
	return result
}
func parser(result *[][]int, level int, node *Node) {
	if node == nil {
		return
	}
	if len((*result)) < level+1 {
		(*result) = append((*result), []int{})
	}
	(*result)[level] = append((*result)[level], node.Val)
	for i := 0; i < len(node.Children); i++ {
		parser(result, level+1, node.Children[i])
	}
}
