## Intuition

The problem statement *is* the algorithm: the maximum is the root, the part to its left
builds the left subtree, the part to its right builds the right subtree. Yesterday split a
list at "the first value bigger than the root"; today the root is the biggest value in the
list and the split is simply around it. The solution writes the definition down as a
recursion.

## Builds on

- [Day 130: Construct Binary Search Tree from Preorder Traversal](https://github.com/architagr/leetcode_solutions/blob/main/medium_problems/1001_1100/construct_binary_search_tree_from_preorder_traversal/) — picking a root from a list and recursing on the pieces either side of it

Like yesterday, each call scans its slice, this time to find the maximum. On a sorted array
the tree is a chain and the scans total O(n²). With 1000 values that's still fast.

There's a well-known O(n) construction, and it's a nice link back to the monotonic stack from
[Day 63: Daily Temperatures](https://github.com/architagr/leetcode_solutions/blob/main/medium_problems/701_800/daily_temperatures/). Walk the
array left to right, keeping a stack of nodes with decreasing values. For each new value, pop
every smaller node; the last one popped becomes the new node's left child (it's the largest
value between the new node and the next bigger one to its left). Then the new node becomes
the right child of whatever is left on top of the stack. The bottom of the stack is the root:

```go
func maxTreeStack(nums []int) *TreeNode {
	stack := []*TreeNode{}
	for _, v := range nums {
		node := &TreeNode{Val: v}
		var last *TreeNode
		for len(stack) > 0 && stack[len(stack)-1].Val < v {
			last = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		}
		node.Left = last
		if len(stack) > 0 {
			stack[len(stack)-1].Right = node
		}
		stack = append(stack, node)
	}
	return stack[0]
}
```

Every node is pushed and popped at most once. I checked it against the recursive version on
two thousand random arrays.

**Complexity (as written):**
- Time: O(n²) worst case (sorted input), O(n log n) when splits are balanced.
- Space: O(h) recursion; the sub-slices share the array.
