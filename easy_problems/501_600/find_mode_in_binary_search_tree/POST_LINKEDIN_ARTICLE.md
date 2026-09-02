## 365 Days of LeetCode Challenge — Day 10/365

# Find Mode in Binary Search Tree

🔗 https://leetcode.com/problems/find-mode-in-binary-search-tree/ · Difficulty: Easy

### The problem

Given the root of a binary search tree (BST) that may contain duplicate values, return
all the mode(s) — the value(s) that occur most frequently. If there's a tie, return all
of them, in any order.

### The intuition

The "BST" in the name is a bit of a red herring for the approach used here. A mode is
just "whichever value shows up most often," and the most direct way to answer that is to
count how many times every value appears, then read off whichever value(s) hit the
highest count. That's a frequency-table problem, and it works exactly the same whether
the tree happens to be sorted or not.

So the approach splits into two clean phases:

1. **Tally every value.** Walk the whole tree once and build a `map[int]int` from value
   → occurrence count. This ignores left/right ordering entirely — every node just
   increments its own value's counter.
2. **Find the max, then collect the winners.** Once every value's count is known, scan
   the map for the highest count, then scan it again to collect every value that hits
   that count — handling ties naturally, since more than one value can be equally
   frequent.

The BST property (`left <= node <= right`) would actually let you solve this with an
in-order traversal and O(1) extra space, by comparing each value to the one right before
it in sorted order — that's exactly what the problem's follow-up question is hinting at.
This implementation trades that space savings for simplicity: a plain hashmap tally is
easy to reason about and get right, at the cost of O(n) extra space for the map.

### The solution

![Example 1](images/1.jpg "Example1")

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

Walking it through `root = [1,null,2,2]` — node `1` at the root with no left child, its
right child is a `2`, and that `2`'s left child is another `2` (expected output `[2]`):

**Phase 1 — `getCnt` tallies every value:**

- `getCnt(1, data)`: `data[1]++` → `data = {1: 1}`.

![Step 1: getCnt visits the root, data becomes {1: 1}](images/walkthrough-1.svg)

- `1.Left` is `nil`, so that branch returns immediately. `getCnt(2a, data)` (the right
  child of `1`): `data[2]++` → `data = {1: 1, 2: 1}`.

![Step 2: getCnt visits 1's right child, data becomes {1: 1, 2: 1}](images/walkthrough-2.svg)

- `getCnt(2b, data)` (the left child of `2a`): `data[2]++` again →
  `data = {1: 1, 2: 2}`. `2b` has no children, so recursion bottoms out.

![Step 3: getCnt visits 2a's left child, data becomes {1: 1, 2: 2}](images/walkthrough-3.svg)

**Phase 2 — `findMode` finds the max, then collects the ties:**

- First pass over `data`: the highest count seen is `cnt = 2`.
- Second pass over `data`: only key `2` has `v == cnt`, so `res = [2]`.

![Step 4: findMode scans for the max count, then collects every key matching it, producing [2]](images/walkthrough-4.svg)

`return res` hands back `[2]`. ✓

**Complexity:** O(n) time — `getCnt` visits every node once, and each scan over `data`
in `findMode` is O(n) in the worst case. O(n) space for the frequency map, plus O(h) for
the recursion stack, where h is the tree's height.

Full code: `easy_problems/501_600/find_mode_in_binary_search_tree/` in the repo.
