## 365 Days of LeetCode Challenge — Day 3/365

# Binary Tree Paths

🔗 https://leetcode.com/problems/binary-tree-paths/ · Difficulty: Easy

### The problem

Given the root of a binary tree, return every root-to-leaf path in any order. A leaf is
a node with no children.

### The intuition

A root-to-leaf path is nothing more than the values you pass through on the way down
from the top to some leaf. So the plan is plain DFS from the root: carry the path built
so far as you descend, and the moment you land on a leaf, that path is done and worth
keeping.

The part that actually needed thought is the root. Every node after it gets appended to
the path as `"->value"`, but the root has nothing before it, no arrow, just its own
value sitting alone. Rather than shove an `if this is the root` check into one big
recursive function, I split it into two: `binaryTreePaths` seeds the path with the
root's value and starts the recursion on its children, and the helper (named `foo`,
which in hindsight I could have called something more useful) handles everything below
the root, where the `"->"` separator always applies.

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

Walking it through the example tree from the problem statement, `root = [1,2,3,null,5]`
(node `2` has a right child `5`, node `3` has no children):

![Example tree](images/1.jpg "Example tree")

- `binaryTreePaths`: `current = "1"`. Recurse into both children.

  ![Walkthrough step 1: current = "1", result = []](images/walkthrough-1.png "Step 1")
- `foo(2, "1", ...)`: `current = "1->2"`. Node `2` has a right child, so recurse into `5`.
- `foo(5, "1->2", ...)`: `current = "1->2->5"`. Leaf. Append `"1->2->5"`.

  ![Walkthrough step 2: current = "1->2->5", result = ["1->2->5"]](images/walkthrough-2.png "Step 2")
- `foo(3, "1", ...)`: `current = "1->3"`. Leaf. Append `"1->3"`.

  ![Walkthrough step 3: current = "1->3", result = ["1->2->5", "1->3"]](images/walkthrough-3.png "Step 3")
- Final result: `["1->2->5", "1->3"]`.

Complexity: O(n²) in the worst case, because building each path string means copying
the prefix so far, and on a skewed tree that copying adds up fast. On a balanced tree
it's closer to O(n log n). Space is O(n) for the recursion stack plus the output.

Full code: `easy_problems/201_300/binary_tree_path/` in the repo.

#DSA #LeetCode #BinaryTree #DFS #Golang #100DaysOfCode #CodingInterview

---

*Solution and code by Archit Agarwal. Write-up drafted with AI assistance from the code and problem statement.*
