**365 Days of LeetCode Challenge — Day 3/365**
**Binary Tree Paths** (Easy)
🔗 https://leetcode.com/problems/binary-tree-paths/

**Intuition**
DFS from the root, building the path as a string while you go, and locking it in the
moment you hit a leaf. The root is the fiddly bit: it's the only node that doesn't get
an arrow before its value, so it needs handling separate from everything below it.

**Full solution**
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

**Walkthrough** on `[1,2,3,null,5]` (node `2` has a right child `5`, node `3` is a leaf):

![Example tree](images/1.jpg "Example tree")

- root `current="1"`, recurse into `2` and `3`

  ![Walkthrough step 1: current = "1", result = []](images/walkthrough-1.svg "Step 1")
- `foo(2,"1")` → `"1->2"`, recurse into `5`
- `foo(5,"1->2")` → `"1->2->5"`, leaf, append

  ![Walkthrough step 2: current = "1->2->5", result = ["1->2->5"]](images/walkthrough-2.svg "Step 2")
- `foo(3,"1")` → `"1->3"`, leaf, append

  ![Walkthrough step 3: current = "1->3", result = ["1->2->5", "1->3"]](images/walkthrough-3.svg "Step 3")
- result: `["1->2->5", "1->3"]`, matches the expected output.

O(n²) worst case on a skewed tree, closer to O(n log n) once the tree is balanced. Not
a case I'd worry about in practice, most trees you actually deal with aren't skewed.
