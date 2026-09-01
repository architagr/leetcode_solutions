## 365 Days of LeetCode Challenge — Day 3/365

# Binary Tree Paths

🔗 https://leetcode.com/problems/binary-tree-paths/ · Difficulty: Easy

### The problem

Given the root of a binary tree, return all root-to-leaf paths in any order. A leaf is
a node with no children.

### The intuition

Every root-to-leaf path is just a sequence of node values from the top of the tree down
to some leaf. The natural way to enumerate all of them is a **DFS from the root**,
carrying along the "path so far" as you descend, and recording that path the moment you
hit a leaf.

There's a small wrinkle: the root itself needs special handling, since the path string
starts as just the root's value with no arrow before it, while every node after that
gets appended as `"->value"`. That's why this solution has two functions instead of
one: `binaryTreePaths` seeds the path with the root's value and kicks off the recursion
into its children, while a helper handles every node after the root, where the `"->"`
separator is always needed.

### The solution

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

Walking it through `root = [1,2,3,null,5]` (node `2` has a right child `5`, node `3`
has no children):
- `binaryTreePaths`: `current = "1"`. Recurse into both children.
- `foo(2, "1", ...)`: `current = "1->2"`. Node `2` has a right child, so recurse into `5`.
- `foo(5, "1->2", ...)`: `current = "1->2->5"`. Leaf — append `"1->2->5"`.
- `foo(3, "1", ...)`: `current = "1->3"`. Leaf — append `"1->3"`.
- Final result: `["1->2->5", "1->3"]`.

**Complexity:** O(n²) worst case (building each path string copies the prefix so far;
across a skewed tree the total work is O(n²)), closer to O(n log n) for a balanced tree.
O(n) space for the recursion stack plus the output.

Full code: `easy_problems/201_300/binary_tree_path/` in the repo.
