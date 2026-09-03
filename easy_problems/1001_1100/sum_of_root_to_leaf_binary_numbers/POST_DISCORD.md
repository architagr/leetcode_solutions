**365 Days of LeetCode Challenge — Day 26/365**
**Sum of Root To Leaf Binary Numbers** (Easy)
🔗 https://leetcode.com/problems/sum-of-root-to-leaf-binary-numbers/

**Intuition:** Each root-to-leaf path spells out a binary number top to bottom, so skip
the collect-then-convert step and carry the running value *down* the recursion
instead. Shift left (`*2`) and drop in the current node's bit at every step. Hit a
leaf and that running value already is the complete binary number for the path, no
extra work needed.

![Example 1](images/1.png "Example1")

**Full solution:**
```go
func sumRootToLeaf(root *TreeNode) int {
	// rootLevelSum starts at 0: no bits accumulated yet above the root.
	return sum(root, 0)
}

// sum returns the total of all root-to-leaf binary numbers in this subtree.
// rootLevelSum is the binary number built from the true root down to (but
// not including) this node, i.e. the bits contributed by ancestors so far.
func sum(root *TreeNode, rootLevelSum int) int {
	if root == nil {
		return 0
	}
	// Shift the accumulated bits one place left to make room for this
	// node's bit, then drop root.Val (0 or 1) into the new units place —
	// the same way a binary number grows digit by digit left to right.
	rootLevelSum *= 2
	currentSum := rootLevelSum + root.Val
	if isLeafNode(root) {
		// currentSum already is the complete binary number for this
		// root-to-leaf path, so it's the whole contribution from here.
		return currentSum
	}

	// Not a leaf: pass the extended currentSum down as the new
	// rootLevelSum for each child, and sum whatever they each contribute.
	leftSum := sum(root.Left, currentSum)
	rightSum := sum(root.Right, currentSum)
	return leftSum + rightSum
}

func isLeafNode(node *TreeNode) bool {
	return node.Left == nil && node.Right == nil
}
```

**Walkthrough** on `[1,0,1,0,1,0,1]` (expected `22`). Quick trace so the shift-and-add
is visible step by step:

- `sum(left=0, 1)` → `sum(leftleft=0, 2)`: `currentSum = 100b = 4`, leaf → `4`

![Step 1: first path root → 0 → 0 reaches a leaf, currentSum = 100b = 4](images/walkthrough-1.svg)

- `sum(leftright=1, 2)`: `currentSum = 101b = 5`, leaf → `5`; `leftSum = 4+5 = 9`

![Step 2: second path root → 0 → 1 reaches a leaf, currentSum = 101b = 5, left subtree total so far = 9](images/walkthrough-2.svg)

- `sum(right=1, 1)` → `sum(rightleft=0, 3)`: `currentSum = 110b = 6`, leaf → `6`

![Step 3: third path root → 1 → 0 reaches a leaf, currentSum = 110b = 6; left subtree already totals 9](images/walkthrough-3.svg)

- `sum(rightright=1, 3)`: `currentSum = 111b = 7`, leaf → `7`; `rightSum = 6+7 = 13`

![Step 4: fourth path root → 1 → 1 reaches a leaf, currentSum = 111b = 7; all four leaves now computed](images/walkthrough-4.svg)

- Back at root: `leftSum(9) + rightSum(13) = 22`, which is the expected total

![Step 5: recursion unwinds, leftSum=9 and rightSum=13 bubble up to root, total = 22](images/walkthrough-5.svg)

O(n) time, since every node is visited once. O(h) space for the call stack, where h is
the tree's height.
