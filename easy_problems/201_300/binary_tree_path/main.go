package binarytreepath

import "fmt"

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func binaryTreePaths(root *TreeNode) []string {
	// The root never gets a "->" prefix, so it's seeded separately from
	// every node below it (handled by foo).
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
	// Every node below the root always needs the "->" separator, unlike
	// the root itself. current is a local copy (strings are immutable in
	// Go), so sibling branches don't interfere with each other's path.
	current += fmt.Sprintf("->%d", node.Val)
	if node.Left == nil && node.Right == nil {
		// Leaf reached: current is now a complete root-to-leaf path.
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
