**365 Days of LeetCode Challenge — Day 12/365**
**Sum of Root To Leaf Binary Numbers** (Easy)
🔗 https://leetcode.com/problems/sum-of-root-to-leaf-binary-numbers/

Each root-to-leaf path spells out a binary number top to bottom, so skip the collect-then-convert step and carry the running value down the recursion instead. Shift left (`*2`) and drop in the current node's bit at every step. Reach a leaf and that running value already is the complete number for the path.

```go
func sumRootToLeaf(root *TreeNode) int {
	return sum(root, 0)
}

func sum(root *TreeNode, rootLevelSum int) int {
	if root == nil {
		return 0
	}
	rootLevelSum *= 2
	currentSum := rootLevelSum + root.Val
	if isLeafNode(root) {
		return currentSum
	}

	leftSum := sum(root.Left, currentSum)
	rightSum := sum(root.Right, currentSum)
	return leftSum + rightSum
}

func isLeafNode(node *TreeNode) bool {
	return node.Left == nil && node.Right == nil
}
```

O(n) time, since every node is visited once. O(h) space for the call stack.

Full walkthrough with step-by-step diagrams: https://github.com/architagr/leetcode_solutions/blob/main/easy_problems/1001_1100/sum_of_root_to_leaf_binary_numbers/SOLUTION.md
