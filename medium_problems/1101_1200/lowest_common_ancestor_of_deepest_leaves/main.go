package lowestcommonancestorofdeepestleaves

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var (
	resFinal    *TreeNode
	deepestPath []*TreeNode
)

func lcaDeepestLeaves(root *TreeNode) *TreeNode {
	resFinal = root
	deepestPath = make([]*TreeNode, 0)
	process(root, []*TreeNode{})
	return resFinal
}

func process(node *TreeNode, path []*TreeNode) {
	if node.Left == nil && node.Right == nil { // this is the leaf node
		if len(deepestPath) > 0 && len(path) == len(deepestPath) {
			resFinal = findLca(path, deepestPath)
		} else if len(path) > len(deepestPath) {
			deepestPath = path
			resFinal = node
		}
	}
	newPath := append([]*TreeNode{}, path...)
	newPath = append(newPath, node)
	if node.Left != nil {
		process(node.Left, newPath)
	}
	if node.Right != nil {
		process(node.Right, newPath)
	}
}
func findLca(list1, list2 []*TreeNode) *TreeNode {
	res := list1[0]
	for i := 1; i < len(list1); i++ {
		if list1[i] == list2[i] {
			res = list1[i]
		}
	}
	return res
}
