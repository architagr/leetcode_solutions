package maximumwidthofbinarytree

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func widthOfBinaryTree(root *TreeNode) int {
	set := make(map[int]int)

	return dfs(root, 1, 0, set)
}

func dfs(node *TreeNode, index, level int, set map[int]int) int {
	if node == nil {
		return 0
	}

	if _, ok := set[level]; !ok {
		set[level] = index
	}

	cur := index - set[level] + 1
	left := dfs(node.Left, index*2-1, level+1, set)
	right := dfs(node.Right, index*2, level+1, set)

	return max(cur, max(left, right))
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
