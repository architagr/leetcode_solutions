## 365 Days of LeetCode Challenge — Day 26/365

# Sum of Root To Leaf Binary Numbers

🔗 https://leetcode.com/problems/sum-of-root-to-leaf-binary-numbers/ · Difficulty: Easy

### The problem

You're given the root of a binary tree where every node holds a `0` or a `1`. Read a
root-to-leaf path top to bottom and its node values spell out a binary number, most
significant bit first — the path `0 -> 1 -> 1 -> 0 -> 1` reads as `01101` in binary,
which is `13`. Sum up that binary number for *every* leaf in the tree and return the
total.

### The intuition

Each root-to-leaf path spells out a binary number one bit at a time, most significant
bit first — exactly the order a top-down recursion visits nodes in. So instead of
collecting the bits into a list and converting them to a number only once a leaf is
reached, the recursion maintains the number *as it descends*: at each step, shift the
value built so far one bit to the left (multiply by 2) and drop in the current node's
bit (`0` or `1`) in the units place. That's just how binary numbers are built digit by
digit — `1101` in binary is `((1*2+1)*2+0)*2+1` — and it maps directly onto "go one
level deeper, append one more bit."

Carrying that running value down as a function parameter means each recursive call
only needs to know "the number formed by the path so far," not the whole path itself.
When a leaf is hit, that running value *is* the complete binary number for that path,
so it's returned directly. Internal nodes don't contribute a value of their own — they
just pass the (now-extended) running value down to their children and sum whatever
comes back up from the left and right subtrees.

### The solution

![Example 1](images/1.png "Example1")

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

Walking it through `root = [1,0,1,0,1,0,1]` (root=1, left=0, right=1, leftleft=0,
leftright=1, rightleft=0, rightright=1), expected `22`:

- `sum(root=1, 0)`: `currentSum = 1` ("bits=1"). Not a leaf, recurse into both
  children.
  - `sum(left=0, 1)`: `currentSum = 10b = 2`. Not a leaf, recurse further.
    - `sum(leftleft=0, 2)`: `currentSum = 100b = 4`. Leaf → returns `4`.

    ![Step 1: first path root → 0 → 0 reaches a leaf, currentSum = 100b = 4](images/walkthrough-1.svg)

    - `sum(leftright=1, 2)`: `currentSum = 101b = 5`. Leaf → returns `5`.
      `leftSum = 4 + 5 = 9`.

    ![Step 2: second path root → 0 → 1 reaches a leaf, currentSum = 101b = 5, left subtree total so far = 9](images/walkthrough-2.svg)

  - `sum(right=1, 1)`: `currentSum = 11b = 3`. Not a leaf, recurse further.
    - `sum(rightleft=0, 3)`: `currentSum = 110b = 6`. Leaf → returns `6`.

    ![Step 3: third path root → 1 → 0 reaches a leaf, currentSum = 110b = 6; left subtree already totals 9](images/walkthrough-3.svg)

    - `sum(rightright=1, 3)`: `currentSum = 111b = 7`. Leaf → returns `7`.
      `rightSum = 6 + 7 = 13`.

    ![Step 4: fourth path root → 1 → 1 reaches a leaf, currentSum = 111b = 7; all four leaves now computed](images/walkthrough-4.svg)

  - Back at the top: `leftSum = 9`, `rightSum = 13`, so the total returned is
    `9 + 13 = 22`. ✓

  ![Step 5: recursion unwinds, leftSum=9 and rightSum=13 bubble up to root, total = 22](images/walkthrough-5.svg)

**Complexity:** O(n) time — every node visited once. O(h) space for the recursion
stack, where h is the tree height (O(log n) balanced, O(n) fully skewed).

Full code: `easy_problems/1001_1100/sum_of_root_to_leaf_binary_numbers/` in the repo.
