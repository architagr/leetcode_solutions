**365 Days of LeetCode Challenge — Day 52/365**
**Validate Binary Search Tree** (Medium)
🔗 https://leetcode.com/problems/validate-binary-search-tree/

The definition sounds local — left smaller, right larger, subtrees also valid — and implementing it literally produces the classic wrong answer: checking each node against its two immediate children. That accepts trees where a node deep in a left subtree exceeds an ancestor several levels up. A BST constrains a node against *every* ancestor, not just its parent.

The way out is a property from Days 29 and 30: in-order traversal of a BST yields ascending values. That's equivalent to the definition, and equivalences run both ways — if the in-order sequence is sorted, it's a BST.

So no bounds get threaded down at all:

```go
func IsValidBst(root *TreeNode) bool {
	arr := InorderTraversal(root)
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] >= arr[i+1] {
			return false
		}
	}
	return true
}
```

On `[5,1,4,null,null,3,6]` the in-order sequence is `[1,5,3,4,6]` — the violating node just lands in the wrong slot, and the adjacent pair `(5,3)` catches it.

`>=` not `>`, because strictly less/greater means duplicates are invalid too.

O(n) time. Comparing against the previous value during the traversal instead of collecting would drop space to O(h).

Full walkthrough with step-by-step diagrams: https://github.com/architagr/leetcode_solutions/blob/main/medium_problems/1_100/validate_binary_search_tree/SOLUTION.md
