---
meta_title: "Path Sum: the leaf check that catches people out"
meta_description: "A node with only a left child is not a leaf. Carry the total down, compare once at the bottom, and test both sides before you decide you are there."
tags: [golang, binary-tree, recursion, dsa, leetcode]
---

# Path Sum

*365 Days of LeetCode Challenge — Day 11/365*

🔗 [LeetCode #112](https://leetcode.com/problems/path-sum/) · Difficulty: Easy

This one looks like it should take two minutes, and it does — unless you get the leaf test wrong,
which is easy to do and produces a solution that passes the obvious examples and fails on a skewed
tree.

The constraint doing all the work is *root-to-leaf*. Not "any path," not "any prefix." A running
total that happens to equal the target halfway down does not count, which rules out the tempting
shortcut of checking at every node and returning early.

### The problem

Given the root of a binary tree and an integer `targetSum`, return `true` if the tree
has a **root-to-leaf** path whose values add up to `targetSum`. A leaf is a node with
no children.

### The intuition

Root-to-leaf is the whole trick here. A match only counts if the path runs all the way
from the root down to a node with no children, so checking "does the running total
equal targetSum" at every node is the wrong approach. An internal node, one that still
has children, can hit the number by coincidence and it still shouldn't count. The
comparison has to happen exactly once, at the leaf.

So the running total gets carried down the tree as an extra argument, picking up each
node's value along the way, and it only gets compared against `targetSum` once you
land on a node with nowhere left to go.

The part I actually had to slow down on: a node with just a left child (no right
child) is not a leaf, even though `root.Right == nil`. It still has somewhere to
recurse into, so scoring it right there would be wrong. That is why the code checks
whether a node has any child to descend into before it ever checks the leaf
condition.

### The solution

![Example 1](https://raw.githubusercontent.com/architagr/leetcode_solutions/main/easy_problems/101_200/path_sum/images/1.jpg)

```go
func hasPathSum(root *TreeNode, targetSum int) bool {
	return sum(root, targetSum, 0)
}

func sum(root *TreeNode, target, current int) bool {
	if root == nil {
		return false
	} else if root.Left != nil && root.Right != nil {
		return sum(root.Left, target, current+root.Val) || sum(root.Right, target, current+root.Val)
	} else if root.Left != nil {
		return sum(root.Left, target, current+root.Val)
	} else if root.Right != nil {
		return sum(root.Right, target, current+root.Val)
	}
	return target == root.Val+current
}
```

Walking it through `root = [5,4,8,11,null,13,4,7,2,null,null,null,1]`, `targetSum = 22`:

- `sum(5, 22, 0)`: `5` has two children, so it recurses left into `4` (`current=5`)
  and, only if that comes back `false`, would try right into `8`.
- `sum(4, 22, 5)`: `4` has only a left child, so it descends into `11`
  (`current=9`).
- `sum(11, 22, 9)`: `11` has two children, so it tries both `7` and `2`, each
  arriving with `current=20`.

![Step 1: current sum descends 5 -> 4 -> 11, current becomes 9](https://raw.githubusercontent.com/architagr/leetcode_solutions/main/easy_problems/101_200/path_sum/images/walkthrough-1.png)

- `sum(7, 22, 20)`: `7` is a leaf. `22 == 7+20` (`27`)? No, so `false`.

![Step 2: at leaf 7, 22 != 27, returns false](https://raw.githubusercontent.com/architagr/leetcode_solutions/main/easy_problems/101_200/path_sum/images/walkthrough-2.png)

- `sum(2, 22, 20)`: `2` is a leaf. `22 == 2+20` (`22`)? Yes, so `true`.

![Step 3: at leaf 2, 22 == 22, returns true](https://raw.githubusercontent.com/architagr/leetcode_solutions/main/easy_problems/101_200/path_sum/images/walkthrough-3.png)

`true` bubbles back up through `11`, then `4`, then `5`. Because Go's `||`
short-circuits, `sum(8, 22, 5)`, the entire right subtree (`13`, `4`, `1`), never
runs at all. That's the part I like about this solution: `hasPathSum` returns `true`
without ever looking at half the tree.

![Step 4: true bubbles up 2 -> 11 -> 4 -> 5, right subtree never visited](https://raw.githubusercontent.com/architagr/leetcode_solutions/main/easy_problems/101_200/path_sum/images/walkthrough-4.png)

**Complexity:** O(n) time, every node is visited at most once (fewer, whenever `||`
short-circuits). O(h) space for the recursion stack, where h is the tree's height.

Full code: `easy_problems/101_200/path_sum/` in the repo.

---

The general lesson is about reading the constraint before reaching for the recursion. "Root-to-leaf"
decides both where the comparison happens and what counts as the bottom, and getting the second one
wrong is invisible until the tree leans one way.

Full code and the step-by-step walkthrough:
[path_sum](https://github.com/architagr/leetcode_solutions/blob/main/easy_problems/101_200/path_sum/SOLUTION.md)

---

*Solution and code by Archit Agarwal. Write-up drafted with AI assistance from the code and problem statement.*
