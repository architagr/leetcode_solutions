# 365 Days of LeetCode Challenge — Day 23/365

## Count Nodes Equal to Average of Subtree

[LeetCode #2265](https://leetcode.com/problems/count-nodes-equal-to-average-of-subtree/) · Medium

![Day 23](HERO.png)

Given the root of a binary tree, count the nodes whose value equals the average of their own subtree, rounded down.

![Example tree](images/1.png)

## Two numbers per node

An average needs a sum and a count. Compute those separately for every node and you're re-walking each subtree once per node, which is O(n^2) on a chain.

But a subtree's sum and size are just its children's sums and sizes plus one node. That means the information can be built once, on the way back up, if the traversal returns it instead of looking it up.

## Answers travel upward

A node can't decide anything on the way down; it hasn't seen what's below it. On the way back up, both children have already reported, and combining is arithmetic:

```
sum   = leftSum + node.Val + rightSum
count = leftCount + rightCount + 1
```

That's post-order. The only structural difference from an ordinary sum-the-tree function is that the recursive call returns a pair.

Once a node has its own sum and count, checking itself is one comparison, and doing that check inside the same call means the running tally comes for free. No second pass.

```go
func averageOfSubtree(root *TreeNode) int {
	var count = 0
	// Returns (sum of subtree, number of nodes in subtree)
	var sumAndCountOfNodes func(node *TreeNode) (currentNodeSum, countNodes int)
	sumAndCountOfNodes = func(node *TreeNode) (currentNodeSum, countNodes int) {
		currentNodeSum, countNodes = 0, 0
		if node == nil {
			return
		}

		// Recursively compute sums and counts for both subtrees
		leftSubtreeNodeSum, leftSubTreeNodeCount := sumAndCountOfNodes(node.Left)
		rightSybTreeNodeSum, rightSubTreeNodeCount := sumAndCountOfNodes(node.Right)

		// Combine subtree values with current node
		currentNodeSum = leftSubtreeNodeSum + node.Val + rightSybTreeNodeSum
		countNodes = leftSubTreeNodeCount + rightSubTreeNodeCount + 1
		// Check if node value equals subtree average
		if currentNodeSum/countNodes == node.Val {
			count++
		}
		return
	}
	sumAndCountOfNodes(root)
	return count
}
```

`count` lives outside the closure and is the answer. The closure captures it, so a match found anywhere increments the same variable. That's why the top-level return value is discarded: the whole tree's sum and count aren't interesting, only the side effect of the checks along the way.

Small Go detail worth knowing: the function is declared with `var` and assigned on the next line, rather than in one statement. A function literal can't refer to its own name inside its initializer, so the two-step form is what makes the recursion compile.

## The nil case is doing real work

`nil` returns `(0, 0)`. It looks like boilerplate, but it's what lets the combine step stay uniform. A node with one child gets `(0, 0)` from the missing side and the formula works with no special case.

That matters more here than in a BST problem, because this is a plain binary tree and one-child nodes are ordinary. In the example, node 5 has only a right child and it's one of the nodes that counts.

## Integer division is load-bearing

The problem says rounded down. Go's `/` on ints truncates toward zero, and with non-negative values truncation is floor, so `currentNodeSum/countNodes` is already correct with no extra work.

Skip past that and you get a wrong answer. Node 5 in the example has sum 11 over 2 nodes. The true average is 5.5, which is not 5. The floored average is 5, which is. Reach for floats and that node silently stops counting.

`countNodes` can't be zero at the division, since the `nil` check returns before it. A node always counts itself.

## Following it through

The example, `root = [4,8,5,0,1,null,6]`:

![Post-order sums and counts](images/walkthrough-1.png)

- **0**, a leaf: sum 0, count 1, average 0. Match.
- **1**, a leaf: sum 1, count 1, average 1. Match.
- **8**: sum `0 + 8 + 1 = 9`, count 3, average 3. Node holds 8, so no match. The only miss in this tree.
- **6**, a leaf: sum 6, count 1, average 6. Match.
- **5**: no left child, so that side reports `(0, 0)`. Sum `0 + 5 + 6 = 11`, count 2, `11 / 2` truncates to 5. Match, and only because of integer division.
- **4**, the root: sum `9 + 4 + 11 = 24`, count 6, average 4. Match.

Five matches. The top-level call returns `(24, 6)` and both are thrown away.

## Cost

**Time O(n).** One visit per node, a few additions and one division each.

**Space O(h)** for the call stack. O(log n) balanced, O(n) for a chain. Nothing allocated beyond the single captured integer.

## Worth noting

The reusable idea here is returning a tuple from a traversal. The moment a node needs more than one fact about its subtree to make a decision, widen the return type rather than adding a second pass. Diameter, balanced-tree checks, and largest-BST-subtree all fall out of the same move.

Full code and the step-by-step walkthrough:
[count_nodes_equal_to_average_of_subtree](https://github.com/architagr/leetcode_solutions/blob/main/medium_problems/2201_2300/count_nodes_equal_to_average_of_subtree/SOLUTION.md)

#DSA #LeetCode #BinaryTree #Recursion #DFS #Golang #CodingInterview

---

*Solution and code by Archit Agarwal. Write-up drafted with AI assistance from the
code and problem statement.*
