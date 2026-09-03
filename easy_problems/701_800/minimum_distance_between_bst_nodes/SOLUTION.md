## Solution walkthrough

The implementation is `minDiffInBST(root *TreeNode) int` in `main.go`, backed by two
helpers, `abs` and `inOrder`.

![Example 1](images/1.jpg "Example1")

We'll trace it on the example above, `root = [4,2,6,1,3]` (expected `1`).

### `minDiffInBST`

1. **Trivial-tree guard.**
   `if root == nil || (root.Left == nil && root.Right == nil) { return 0 }` — if the
   tree is empty, or is a single node with no children at all, there's no pair of nodes
   to compare, so there's nothing to return but `0`. (The problem's own constraints
   guarantee at least 2 nodes, so in practice this mainly guards `root == nil`.)

2. **Flatten the tree into a sorted list.** `arr := inOrder(root)` — because `root` is a
   BST, visiting it in-order (left subtree, node, right subtree) is guaranteed to visit
   every value in strictly increasing order. That turns "compare every pair of nodes"
   into "sort once, then only compare neighbors."

   ![Step 1: in-order traversal collects [1, 2, 3, 4, 6] from the tree](images/walkthrough-1.png)

3. **Scan for the smallest neighboring gap.**
   ```go
   minVal := math.MaxInt
   for i := 1; i < len(arr); i++ {
       minVal = min(minVal, abs(arr[i]-arr[i-1]))
   }
   ```
   `minVal` starts at the largest possible `int` so that the very first comparison
   always wins. The loop then walks `arr` one adjacent pair at a time — `(1,2)`,
   `(2,3)`, `(3,4)`, `(4,6)` — taking the absolute difference of each pair and keeping
   the smallest one seen so far. Because the minimum difference between *any* two BST
   node values can only occur between two values that are adjacent once sorted, this
   single linear pass is enough; no need to check non-adjacent pairs.

   ![Step 2: scanning adjacent gaps 1, 1, 1, 2 — minVal settles at 1](images/walkthrough-2.png)

4. **Return the answer.** `return minVal` — after the loop, `minVal` holds the smallest
   gap found across the whole sorted sequence.

   ![Step 3: minDiffInBST(root) returns 1](images/walkthrough-3.png)

### `abs`

```go
func abs(a int) int {
	if a < 0 {
		return -1 * a
	}
	return a
}
```
A small hand-rolled absolute value: negative inputs get flipped, non-negative inputs
pass through unchanged. Used to turn `arr[i]-arr[i-1]` (always non-negative here, since
`arr` is sorted, but written generically) into a plain magnitude.

### `inOrder`

```go
func inOrder(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	return append(inOrder(root.Left), append([]int{root.Val}, inOrder(root.Right)...)...)
}
```
The base case returns an empty slice for a `nil` subtree. Otherwise it recursively
collects the left subtree's values, then this node's own `Val`, then the right
subtree's values, and concatenates them in that order via two nested `append` calls —
exactly the "left, node, right" recipe that makes in-order traversal of a BST come out
sorted.

**Complexity:** the traversal visits each of the `n` nodes once, but because
`inOrder` rebuilds and copies a slice at every recursive call via nested `append`s, the
total copying work can degrade toward O(n²) on a heavily skewed tree rather than staying
linear (a straightforward accumulator-based traversal would avoid this). The final scan
over the sorted array is a clean O(n). Space is O(n) for the collected slice plus O(h)
for the recursion stack, where `h` is the tree's height.
