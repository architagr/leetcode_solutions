**365 Days of LeetCode Challenge — Day 41/365**
**Find a Corresponding Node of a Binary Tree in a Clone of That Tree** (Easy)
🔗 https://leetcode.com/problems/find-a-corresponding-node-of-a-binary-tree-in-a-clone-of-that-tree/

`cloned` is a structurally identical copy of `original`, so you never have to search it on its own. Walk both trees in lockstep, one recursive call advancing both pointers, and compare values at each paired position. The instant they match, the `cloned`-side pointer is standing exactly where `target` stands.

```go
func getTargetCopy(original, clones, target *TreeNode) *TreeNode {
	if clones == nil {
		return nil
	}
	if clones.Val == target.Val {
		return clones
	}
	n := getTargetCopy(original.Left, clones.Left, target)
	if n != nil {
		return n
	}
	return getTargetCopy(original.Right, clones.Right, target)
}
```

O(n) time, O(h) space.

Full walkthrough: https://github.com/architagr/leetcode_solutions/blob/main/easy_problems/1301_1400/find_a_corresponding_node_of_a_binary_tree_in_a_clone_of_that_tree/SOLUTION.md
