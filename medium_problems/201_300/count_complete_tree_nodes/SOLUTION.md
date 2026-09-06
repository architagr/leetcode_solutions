# Solution Walkthrough

## Approach: Recursive Traversal

This solution uses a simple recursive approach: to count nodes in a tree, count the nodes in the right subtree, count the nodes in the left subtree, and add 1 for the root.

## Code Breakdown

```go
func countNodes(root *TreeNode) int {
    if root == nil {
        return 0
    }
    
    rightCount := countNodes(root.Right)
    leftCount := countNodes(root.Left)
    return rightCount + leftCount + 1
}
```

**Base case:** If we reach a nil node, there are no nodes to count, so return 0.

**Recursive case:** 
1. Count all nodes in the right subtree by recursively calling `countNodes(root.Right)`
2. Count all nodes in the left subtree by recursively calling `countNodes(root.Left)`
3. Add 1 for the current node
4. Return the total

## Example Trace

For the tree `[1,2,3,4,5,6]`:

```
        1
       / \
      2   3
     / \
    4   5
```

We start at node 1:
- Recursively count the right subtree (node 3 and below): returns 1
- Recursively count the left subtree (node 2, 4, 5): returns 3
  - Which recursively counts its children and their subtrees
- Add 1 for the root node
- Total: 1 + 3 + 1 = 5... wait that's wrong. Let me retrace.

Actually:
- countNodes(1):
  - rightCount = countNodes(3) = 1 (just node 3, no children)
  - leftCount = countNodes(2) 
    - rightCount = countNodes(5) = 1
    - leftCount = countNodes(4) = 1
    - return 1 + 1 + 1 = 3
  - return 3 + 1 + 1 = 5

Hmm, but there are 6 nodes total. Let me recount the tree:
```
        1
       / \
      2   3
     / \
    4   5  6?
```

No wait, the example says `[1,2,3,4,5,6]` which is 6 nodes total. The tree layout should be:
```
        1
       / \
      2   3
     / \ /
    4  5 6
```

Let me retrace:
- countNodes(1):
  - rightCount = countNodes(3):
    - rightCount = countNodes(nil) = 0
    - leftCount = countNodes(6) = 1
    - return 0 + 1 + 1 = 2
  - leftCount = countNodes(2):
    - rightCount = countNodes(5) = 1
    - leftCount = countNodes(4) = 1
    - return 1 + 1 + 1 = 3
  - return 2 + 3 + 1 = 6 ✓

## Complexity

- **Time:** O(n) — we visit every node once
- **Space:** O(h) where h is the height — recursion stack depth
