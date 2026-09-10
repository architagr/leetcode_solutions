**365 Days of LeetCode Challenge — Day 49/365**
**Binary Tree Upside Down** (Medium)
🔗 https://leetcode.com/problems/binary-tree-upside-down/

Thirteen lines doing two unrelated jobs on one recursion, and separating them is most of understanding it.

Job one: find the new root — the deepest node on the left spine — and pass it up. That's the base case, and it's the only place a value is ever produced. Every frame above just returns `x` untouched.

Job two: local pointer surgery at each level, on the way back up. `root.Left` is promoted above `root`, the old right child becomes its left, the old parent becomes its right. After the recursive call, because the subtree has to be flipped before its old parent can hang underneath it.

```go
x := upsideDownBinaryTree(root.Left)
newRight := root
newNode := root.Left
newLeft := root.Right
newNode.Left = newLeft
newNode.Right = newRight
root.Left = nil
root.Right = nil
return x
```

Those two nils look like they'd destroy the work. They don't: for every node except the original root, the parent's frame overwrites both pointers moments later. The clearing only survives at the top, where the old root really does become a leaf.

O(h) time — the recursion never descends right.

Full walkthrough with step-by-step diagrams: https://github.com/architagr/leetcode_solutions/blob/main/medium_problems/101_200/binary_tree_upside_down/SOLUTION.md
