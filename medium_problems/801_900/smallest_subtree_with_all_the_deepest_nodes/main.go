package smallestsubtreewithallthedeepestnodes

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var (
	data     map[int][][]*TreeNode
	maxDepth int
)

func subtreeWithAllDeepest(root *TreeNode) *TreeNode {
	data = make(map[int][][]*TreeNode)
	maxDepth = -1
	dfs(root, []*TreeNode{})
	l := data[maxDepth]
	if len(l) == 1 {
		return l[0][len(l[0])-1]
	}
	res := root
	for i := 0; i < maxDepth; i++ {
		x := l[0][i]
		for j := 0; j < len(l); j++ {
			if l[j][i] != x {
				x = nil
				break
			}
		}
		if x != nil {
			res = x
		}
	}
	return res

}

func dfs(node *TreeNode, path []*TreeNode) {
	path = append(path, node)
	if node.Left == nil && node.Right == nil {
		maxDepth = maxVal(maxDepth, len(path))
		if len(path) == maxDepth {
			c := make([]*TreeNode, len(path))
			copy(c, path)
			data[len(path)] = append(data[len(path)], c)
		}

	} else {
		if node.Left != nil {
			dfs(node.Left, path)
		}
		if node.Right != nil {
			dfs(node.Right, path)
		}
	}
	path = path[:len(path)-1]
}

func maxVal(a, b int) int {
	if a > b {
		return a
	}
	return b
}
