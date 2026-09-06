# Solution walkthrough

The code is two functions. `bstToGst` is a one-line wrapper; `parse` does the work.

```go
func bstToGst(root *TreeNode) *TreeNode {
	parse(root, 0)
	return root
}
```

The `0` is the seed for the running sum: before anything is visited, nothing greater has been seen. The tree is modified in place, so the same `root` pointer comes back out.

## The recursive step

```go
func parse(node *TreeNode, parentSum int) int {
	if node == nil {
		return parentSum + 0
	}
	// Traverse right subtree first to process larger values
	right := parse(node.Right, parentSum)
	// Add the accumulated sum of all greater nodes
	node.Val += right
	// Traverse left subtree with updated sum
	return parse(node.Left, node.Val)
}
```

The contract for `parse` is: given the sum of every key already processed, handle this subtree and return the sum of every key processed once that subtree is done.

The `nil` case returns `parentSum` untouched, so an empty branch contributes nothing and the total flows past it. (`parentSum + 0` is just an explicit way of writing that.)

Then three lines in a deliberate order:

1. `parse(node.Right, parentSum)` runs first. Everything in the right subtree is greater than `node`, so those keys must be folded in before `node` is touched. The incoming `parentSum` passes down unchanged.
2. `node.Val += right` converts this node. At this moment `right` holds the sum of every key greater than `node`, so the addition is exactly the definition of the problem.
3. `parse(node.Left, node.Val)` recurses left with `node.Val` as the new running sum. This works because `node.Val` was just updated to equal the node's key plus everything above it, which is precisely the total the left subtree needs.

That last point is the neat part. There is no separate accumulator variable, because after step 2 the node itself holds the running sum.

## Tracing it

Take the BST `[4,1,6,0,2,5,7]`:

![Visit order and the running sum](images/walkthrough-1.png)

The orange badges give the order `parse` reaches each node: 7, 6, 5, 4, 2, 1, 0. Strictly descending, which is the whole point of going right before left.

Following the table one row at a time:

- **7** is the rightmost node. `parentSum` is 0, so it stays 7. Returns 7.
- **6** receives 7 from its right child. `6 + 7 = 13`. Returns 13 into its left child.
- **5** is 6's left child and receives 13. `5 + 13 = 18`. Returns 18, which unwinds all the way back to the root's `right` variable.
- **4** is the root. `right` is 18, so `4 + 18 = 22`. Correct: the keys above 4 are 5, 6 and 7, which sum to 18.
- **2** gets 22 from the root's left call. `2 + 22 = 24`.
- **1** gets 24. `1 + 24 = 25`.
- **0** gets 25. `0 + 25 = 25`. Nothing left to visit.

The finished tree:

![The converted tree](images/walkthrough-2.png)

Spot-check node 5: the keys greater than it are 6 and 7, and `5 + 6 + 7 = 18`. Node 7 never changes, because nothing in the tree is bigger than it.

## Complexity

**Time: O(n).** Every node is visited exactly once and does a constant amount of work.

**Space: O(h)** for the call stack, h being the tree height. Balanced comes out to O(log n), a degenerate one-sided tree to O(n). Nothing is allocated beyond the recursion itself, and the conversion is in place.
