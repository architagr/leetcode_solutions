**365 Days of LeetCode Challenge — Day 10/365**
**Find Mode in Binary Search Tree** (Easy)
🔗 https://leetcode.com/problems/find-mode-in-binary-search-tree/

**Intuition:** The "BST" part is a bit of a red herring here — the simplest approach
ignores the ordering entirely and just tallies how many times every value occurs in a
hashmap, then reads off whichever value(s) hit the highest count (handling ties, since
there can be more than one mode).

![Example 1](images/1.jpg "Example1")

**Full solution:**
```go
func findMode(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	// Tally how many times every value occurs. This ignores the BST
	// ordering entirely and treats the tree like any binary tree.
	data := make(map[int]int)
	data = getCnt(root, data)

	res := make([]int, 0)
	cnt := 0
	// First pass: find the highest frequency present in the tree.
	for _, v := range data {
		if v > cnt {
			cnt = v
		}
	}

	// Second pass: collect every value whose frequency matches the max,
	// so ties (multiple modes) are all returned, not just one.
	for k, v := range data {
		if v == cnt {
			res = append(res, k)
		}
	}
	return res
}

// getCnt walks the whole tree and increments data[node.Val] for every
// node visited. Since data is a map (reference type), every recursive
// call mutates the same underlying map, so the return value only
// matters to the top-level caller.
func getCnt(root *TreeNode, data map[int]int) map[int]int {
	if root == nil {
		return nil
	}
	data[root.Val]++
	getCnt(root.Left, data)
	getCnt(root.Right, data)
	return data
}
```

**Walkthrough** on `[1,null,2,2]` (`1` at the root, right child `2`, that `2`'s left
child another `2`; expected `[2]`):

- `getCnt(1, data)` → `data = {1: 1}`

![Step 1: getCnt visits the root, data becomes {1: 1}](images/walkthrough-1.svg)

- `getCnt(2a, data)` (right child of `1`) → `data = {1: 1, 2: 1}`

![Step 2: getCnt visits 1's right child, data becomes {1: 1, 2: 1}](images/walkthrough-2.svg)

- `getCnt(2b, data)` (left child of `2a`) → `data = {1: 1, 2: 2}`

![Step 3: getCnt visits 2a's left child, data becomes {1: 1, 2: 2}](images/walkthrough-3.svg)

- `findMode` scans for the max (`cnt = 2`), then collects every key matching it → `[2]`

![Step 4: findMode scans for the max count, then collects every key matching it, producing [2]](images/walkthrough-4.svg)

O(n) time, O(n) space for the map (plus O(h) recursion stack).
