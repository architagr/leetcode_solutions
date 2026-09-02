**365 Days of LeetCode Challenge — Day 28/365**
**Evaluate Boolean Binary Tree** (Easy)
🔗 https://leetcode.com/problems/evaluate-boolean-binary-tree/

**Intuition:** The tree itself *is* the boolean expression — leaves hold `True`/`False`
literals, internal nodes hold `OR`/`AND` operators. Post-order recursion evaluates both
children before combining them with the parent's operator, exactly like evaluating a
nested expression from the inside out.

![Example 1](images/1.png "Example1")

**Full solution:**
```go
const (
	FALSE = 0
	TRUE  = 1
	OR    = 2
	AND   = 3
)

func evaluateTree(root *TreeNode) bool {
	if root.Left == nil && root.Right == nil {
		return root.Val == TRUE
	}

	left := evaluateTree(root.Left)
	right := evaluateTree(root.Right)
	if root.Val == OR {
		return left || right
	}
	return left && right
}
```

**Walkthrough** on `[2,1,3,null,null,0,1]` (expected `true`):
- Root is `OR`, left child is leaf `1`, right child is `AND` with leaves `0` and `1`.
- The three leaves hit the base case: root's left child → `True`, `AND`'s left child →
  `False`, `AND`'s right child → `True`.

![Step 1: leaves hit the base case, root.Val == TRUE decides each](images/walkthrough-1.svg)

- `AND` node combines `False && True = False`.

![Step 2: AND node combines left && right = False && True = False](images/walkthrough-2.svg)

- Root `OR` combines `True || False = True` ✓

![Step 3: root OR combines left || right = True || False = True](images/walkthrough-3.svg)

O(n) time, O(h) space (tree height).
