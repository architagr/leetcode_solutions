**365 Days of LeetCode Challenge — Day 39/365**
**Delete Node in a BST** (Medium)
🔗 https://leetcode.com/problems/delete-node-in-a-bst/

The problem splits itself: find the node, then delete it. Finding it is Day 26's walk — one comparison picks the only direction it could be in.

Deleting is the interesting half, because a node with two children can't just be unhooked. Something has to fill the hole, and only two values in the entire tree can: the in-order predecessor and successor.

That follows straight from Day 32's property. If in-order has to stay sorted, the replacement must sit between everything in the left subtree and everything in the right — so it's the largest value on the left or the smallest on the right. Nothing else works.

So the code never removes an internal node at all:

```go
root.Val = successor(root)
root.Right = deleteNode(root.Right, root.Val)
```

Overwrite the value, then delete the duplicate below. That looks circular but terminates: the successor is the leftmost node of the right subtree, so it has no left child, so each step recurses into a strictly easier case.

Also note `root.Left = deleteNode(...)` rather than discarding the result — deletion can change which node roots a subtree, so callers have to re-attach.

O(h) time, O(h) space.

Full walkthrough with step-by-step diagrams: https://github.com/architagr/leetcode_solutions/blob/main/medium_problems/401_500/delete_node_in_a_bst/SOLUTION.md
