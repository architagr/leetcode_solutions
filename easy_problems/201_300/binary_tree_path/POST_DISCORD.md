**365 Days of LeetCode Challenge — Day 3/365**
**Binary Tree Paths** (Easy)
🔗 https://leetcode.com/problems/binary-tree-paths/

DFS from the root, building the path as a string while you go, and locking it in the moment you hit a leaf. The root is the fiddly bit: it's the only node that doesn't get an arrow before its value, so it needs handling separate from everything below it.

```go
func binaryTreePaths(root *TreeNode) []string {
	current := fmt.Sprint(root.Val)
	if root.Left == nil && root.Right == nil {
		return []string{current}
	}
	result := make([]string, 0)
	if root.Left != nil {
		foo(root.Left, current, &result)
	}
	if root.Right != nil {
		foo(root.Right, current, &result)
	}
	return result
}

func foo(node *TreeNode, current string, result *[]string) {
	current += fmt.Sprintf("->%d", node.Val)
	if node.Left == nil && node.Right == nil {
		*result = append(*result, current)
		return
	}
	if node.Left != nil {
		foo(node.Left, current, result)
	}
	if node.Right != nil {
		foo(node.Right, current, result)
	}
}
```

Full walkthrough with step-by-step diagrams: https://github.com/architagr/leetcode_solutions/blob/main/easy_problems/201_300/binary_tree_path/SOLUTION.md
