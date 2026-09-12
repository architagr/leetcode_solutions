**365 Days of LeetCode Challenge — Day 17/365**
**Construct String from Binary Tree** (Medium)
🔗 https://leetcode.com/problems/construct-string-from-binary-tree/

The traversal is plain pre-order. The medium is one clause in the formatting rules: empty parens are omitted, except when a node has a right child and no left child, where you must emit `()` anyway.

That exception isn't decoration. Without it `1(2)` would describe both a left child and a right child, and the representation stops mapping one-to-one onto the tree.

The trick is to make the code asymmetric the way the rules are. Once a node is known to have any child, emit the left pair unconditionally — `parse(nil)` returns `""`, so that same line produces the `()` placeholder for free. No branch anywhere tests for "right but no left".

```go
func parse(node *TreeNode) string {
	if node == nil {
		return ""
	}
	s := strconv.Itoa(node.Val)
	if node.Right != nil || node.Left != nil {
		s += "(" + parse(node.Left) + ")"
		if node.Right != nil {
			s += "(" + parse(node.Right) + ")"
		}
	}
	return s
}
```

O(n) visits, though `+=` on strings makes the copying O(n²) on a skewed tree — a strings.Builder would fix that.

Full walkthrough with step-by-step diagrams: https://github.com/architagr/leetcode_solutions/blob/main/medium_problems/601_700/construct_string_from_binary_tree/SOLUTION.md
