---
meta_title: "The k closest values in a sorted array are contiguous"
meta_description: "That one claim turns pick-k-things into pick-a-window. Flatten the BST in-order, find the closest value, and expand outward taking the nearer neighbour."
tags: [golang, binary-search-tree, two-pointers, in-order, leetcode]
---

# Closest Binary Search Tree Value II

*365 Days of LeetCode Challenge — Day 34/365*

🔗 [LeetCode #272](https://leetcode.com/problems/closest-binary-search-tree-value-ii/) · Difficulty: Hard

Both hards in this batch have had the same shape: a traversal you already know, plus one
observation that has to be made before any of it helps.

Day 25's observation was that a node computes two different values. Today's is a claim about
sorted arrays that has nothing to do with trees at all, and the tree part of the problem
becomes routine the moment you have it.

### The problem

Given the root of a BST, a target value, and an integer `k`, return the `k` values closest
to the target, in any order. You're guaranteed a unique answer.

![Example 1](images/1.jpg)

### The intuition

The second hard of the challenge, and like the first one it isn't hard because of the
traversal. It's hard because of a claim you have to notice and trust.

Flatten the BST in-order and you get the values in ascending order — the property this
batch has now used five times, and yesterday's iterator did exactly this flattening.

Now the claim: in a sorted array, the k values closest to a target are always contiguous.

That's worth a second, because it's what turns the problem from "pick k things" into "pick
a window". Suppose the answer skipped some value `v` and instead included a value `w`
further from the target, with `v` sitting between the closest element and `w`. Because the
array is sorted, `v` lies between them in value too, so `v` is at least as close to the
target as `w` is. Swapping `w` for `v` gives an answer at least as good. A non-contiguous
answer is therefore never strictly better than the contiguous one covering the same span.

Once that's settled the algorithm falls out. Find the single closest value, then grow
outward from it, always taking whichever neighbour is nearer the target. Stop at k. Greedy
is safe because distance increases monotonically as you move away from the target in either
direction.

Two pointers walking outward from a centre, repeatedly taking the smaller of two
candidates, is the merge step of a merge sort read backwards — the code's own comment says
as much.

### Builds on

- [Day 33: Binary Search Tree Iterator](https://github.com/architagr/leetcode_solutions/blob/main/medium_problems/101_200/binary_search_tree_iterator/) — flattening a BST in-order into a slice, which is the first half of this solution verbatim

### The solution

```go
func closestKValues(root *TreeNode, target float64, k int) []int {
	min = math.MaxInt64
	minIndex = -1
	inOrderArr := make([]int, 0)
	inOrderTraversal(root, target, &inOrderArr)
	l := 0
	l++
	i, j := minIndex-1, minIndex+1
	for i >= 0 && j < len(inOrderArr) && l < k {
		if absDiff(inOrderArr[i], target) < absDiff(inOrderArr[j], target) {
			i--
		} else {
			j++
		}
		l++
	}
	for ; i >= 0 && l < k; l++ {
		i--
	}
	for ; j < len(inOrderArr) && l < k; l++ {
		j++
	}
	return inOrderArr[i+1 : j]
}
```

Tracing `root = [4,2,5,1,3]`, `target = 3.714286`, `k = 2`. In-order gives `[1,2,3,4,5]`,
and the expected answer is `[4,3]`.

One traversal does two jobs: it flattens the tree and tracks the smallest difference seen
so far. `minIndex` is captured as `len(*arr)` *before* the append, which is precisely the
index the value is about to occupy — so no separate search for the closest element is
needed.

![Step 1: the flatten computes diffs as it goes](images/walkthrough-1.png)

The diffs come out `2.714, 1.714, 0.714, 0.286, 1.286`, so `minIndex` is `3`.

![Step 2: minIndex settles on the single closest value](images/walkthrough-2.png)

`l := 0; l++` is a roundabout way of writing `l := 1`, and the 1 is right — the closest
element is already part of the answer, so only `k-1` more are needed.

![Step 3: the left neighbour is nearer, so i moves](images/walkthrough-3.png)

The two trailing loops handle running out of array: if one side is exhausted before `k`
values are collected, the other supplies the rest. Only one of them can ever execute.

Both pointers end up one step outside the collected range, so `i+1` and `j` are exactly the
inclusive-exclusive bounds of the window.

![Step 4: the window between the pointers is the answer](images/walkthrough-4.png)

One thing worth flagging about the implementation: `min` and `minIndex` are package-level
variables rather than parameters. It works, because the function resets both at the top of
every call, but it makes the function non-reentrant — two goroutines calling it
concurrently would corrupt each other's search. Threading them through as pointers, the way
`arr` already is, would remove that.

O(n) for the traversal plus O(k) for the expansion. Space is O(n) for the flattened slice
plus O(h) for the recursion stack. The known better answer is O(k + h): run two stack-based
BST iterators outward from the target — yesterday's follow-up applied twice, one backwards
and one forwards — and merge `k` values off them.

---

It is worth practising the habit of asking what shape the answer has before asking how to
compute it. "The answer is some k elements" and "the answer is a contiguous run of k
elements" are very different problems, and only one of them is solved by two pointers.

The proof is short enough to reconstruct at a whiteboard, which is the real reason to know
it rather than to remember the technique.

Full code and the step-by-step walkthrough:
[closest_binary_search_tree_value_ii](https://github.com/architagr/leetcode_solutions/blob/main/hard_problems/201_300/closest_binary_search_tree_value_ii/SOLUTION.md)

---

*Solution and code by Archit Agarwal. Write-up drafted with AI assistance from the code and problem statement.*
