---
meta_title: "Level averages without BFS: index by depth"
meta_description: "DFS does not visit a tree level by level, and it does not need to. Key the running sum and count by depth and the visit order stops mattering."
tags: [golang, binary-tree, dfs, recursion, leetcode]
---

# Average of Levels in Binary Tree

*365 Days of LeetCode Challenge — Day 14/365*

🔗 [LeetCode #637](https://leetcode.com/problems/average-of-levels-in-binary-tree/) · Difficulty: Easy

"Per level" almost always means BFS — a queue, drain one level, move to the next. That solution is
fine and it's what most people write here.

This one takes the other road on purpose, because the reason it works is more interesting than the
problem. DFS visits a tree branch by branch, in an order that has nothing to do with levels. It
still computes correct per-level averages, and the trick is that it never needs to know what order
the nodes arrived in.

### The problem

Given a binary tree's root, return the average node value at each level, ordered from
the root down.

### The intuition

Most people reach for BFS here: walk the tree one level at a time with a queue,
average each level as you drain it, then move to the next. I went a different way and
used **depth-first search** instead, and it still lands on the same per-level
averages, by carrying the current depth as a parameter and using it as an index into
two running-total arrays.

The part that took me a second to see: DFS doesn't visit nodes level by level, it
visits them branch by branch. So instead of finishing one level before starting the
next, the code keeps a running sum and count per level, stored as `result[level]` and
`count[level]`. Every visit adds its value into the slot for its own depth, no matter
which branch it came down. A node that reaches depth 2 first and one that gets there
last both end up in `result[2]` and `count[2]`, because the slot is keyed by level, not
by when the recursion happens to arrive there.

The two arrays start empty and grow as needed. The first time the recursion hits a new
depth, it appends a fresh slot to both `result` and `count`. Every later visit at that
depth just reuses the slot instead of appending again. Once the whole tree's been
walked, one final pass divides each level's sum by its count and turns the running
totals into averages.

### The solution

![Example 1](https://raw.githubusercontent.com/architagr/leetcode_solutions/main/easy_problems/601_700/average_of_levels_in_binary_tree/images/1.jpg "Example1")

```go
func AverageOfLevel(root *TreeNode) []float64 {
	result := make([]float64, 0)
	count := make([]float64, 0)
	getSumAndCount(root, &result, &count, 0)
	for i := range result {
		result[i] /= count[i]
	}
	return result
}

func getSumAndCount(node *TreeNode, result, count *[]float64, level int) {
	if node == nil {
		return
	}
	if len(*result) < level+1 {
		*result = append(*result, 0.0)
		*count = append(*count, 0.0)
	}

	(*result)[level] += float64(node.Val)
	(*count)[level]++
	getSumAndCount(node.Left, result, count, level+1)
	getSumAndCount(node.Right, result, count, level+1)
}
```

Here's the walkthrough on `[3,9,20,null,null,15,7]` (expected `[3.0, 14.5, 11.0]`):

- Visit `3` (level 0): no slot yet, so one gets created: `result=[3.0]`,
  `count=[1.0]`.

  ![Step 1: visit 3 (level 0), create slot 0, result[0]=3, count[0]=1](https://raw.githubusercontent.com/architagr/leetcode_solutions/main/easy_problems/601_700/average_of_levels_in_binary_tree/images/walkthrough-1.png)

- Visit `9` (level 1): slot 1 doesn't exist yet either, so it gets created too:
  `result=[3.0, 9.0]`, `count=[1.0, 1.0]`.

  ![Step 2: visit 9 (level 1), create slot 1, result[1]=9, count[1]=1](https://raw.githubusercontent.com/architagr/leetcode_solutions/main/easy_problems/601_700/average_of_levels_in_binary_tree/images/walkthrough-2.png)

- Visit `20` (level 1): slot 1 is already there from the `9` visit, so this one
  updates it in place instead of appending: `result[1] += 20` → `29.0`, `count[1]++`
  → `2.0`. `20` sits in a completely different branch than `9`, but they both write
  into the same slot because they're at the same depth. That's the whole trick.

  ![Step 3: visit 20 (level 1), slot 1 exists, result[1] becomes 29, count[1] becomes 2](https://raw.githubusercontent.com/architagr/leetcode_solutions/main/easy_problems/601_700/average_of_levels_in_binary_tree/images/walkthrough-3.png)

- Visit `15` (level 2): no slot 2 yet, so it's created: `result=[3.0, 29.0, 15.0]`,
  `count=[1.0, 2.0, 1.0]`.

  ![Step 4: visit 15 (level 2), create slot 2, result[2]=15, count[2]=1](https://raw.githubusercontent.com/architagr/leetcode_solutions/main/easy_problems/601_700/average_of_levels_in_binary_tree/images/walkthrough-4.png)

- Visit `7` (level 2): slot 2 already exists, updated in place: `result[2] += 7` →
  `22.0`, `count[2]++` → `2.0`.

  ![Step 5: visit 7 (level 2), slot 2 exists, result[2] becomes 22, count[2] becomes 2](https://raw.githubusercontent.com/architagr/leetcode_solutions/main/easy_problems/601_700/average_of_levels_in_binary_tree/images/walkthrough-5.png)

- Recursion's done: `result=[3.0, 29.0, 22.0]`, `count=[1.0, 2.0, 2.0]`. The final
  loop divides each slot: `3.0/1.0 = 3.0`, `29.0/2.0 = 14.5`, `22.0/2.0 = 11.0`, giving
  `[3.0, 14.5, 11.0]`, which matches.

  ![Step 6: divide result[i] by count[i] for every level, giving 3.0, 14.5, 11.0](https://raw.githubusercontent.com/architagr/leetcode_solutions/main/easy_problems/601_700/average_of_levels_in_binary_tree/images/walkthrough-6.png)

**Complexity:** O(n) time, since every node gets visited exactly once and the final
division pass only runs once per level. Space is O(h) for the recursion stack (h is
tree height), plus O(L) for the `result`/`count` slices, where L (the number of
levels) tops out at h + 1.

---

The transferable idea: when accumulation is keyed by a property of the item rather than by when it
showed up, traversal order becomes free. That's what makes the "wrong" traversal work here, and
it's worth recognising the next time a problem seems to demand a specific order.

Full code and the step-by-step walkthrough:
[average_of_levels_in_binary_tree](https://github.com/architagr/leetcode_solutions/blob/main/easy_problems/601_700/average_of_levels_in_binary_tree/SOLUTION.md)

---

*Solution and code by Archit Agarwal. Write-up drafted with AI assistance from the code and problem statement.*
