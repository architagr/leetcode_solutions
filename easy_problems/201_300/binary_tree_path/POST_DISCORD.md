**365 Days of LeetCode Challenge — Day 3/365**
**Binary Tree Paths** (Easy)
🔗 https://leetcode.com/problems/binary-tree-paths/

**Intuition:** DFS from the root, carrying the path-so-far as a string, recording it
when you hit a leaf. Root needs special handling — it's the only node with no `"->"`
before its value.

**Full solution:**
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

**Walkthrough** on `[1,2,3,null,5]`:
- root `current="1"`, recurse into `2` and `3`
- `foo(2,"1")` → `"1->2"`, recurse into `5`
- `foo(5,"1->2")` → `"1->2->5"`, leaf → append
- `foo(3,"1")` → `"1->3"`, leaf → append
- **result:** `["1->2->5", "1->3"]` ✓

O(n²) worst case (skewed tree), O(n log n) balanced.
