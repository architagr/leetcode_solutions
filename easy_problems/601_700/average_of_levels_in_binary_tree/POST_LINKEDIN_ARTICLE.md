## 365 Days of LeetCode Challenge — Day 16/365

# Average of Levels in Binary Tree

🔗 https://leetcode.com/problems/average-of-levels-in-binary-tree/ · Difficulty: Easy

### The problem

Given the root of a binary tree, return the average value of the nodes on each level,
as an array ordered from the root's level downward.

### The intuition

The obvious way to solve this is breadth-first search: drain the tree one level at a
time with a queue, average each level as you finish it, then move on. This solution
takes a different route — a **depth-first search** that still lands on the same
per-level averages, by carrying the current depth as a parameter and using it as an
index into two running-total arrays.

The trick is that a DFS doesn't visit nodes level by level, it visits them branch by
branch. So instead of finishing one level before starting the next, the solution keeps
a running sum and count **per level**, stored as `result[level]` and `count[level]`,
and lets every visit — no matter which branch of the tree it comes from — add its value
into the slot for its own depth. Whichever node happens to reach depth 2 first, second,
or last, they all land in `result[2]` and `count[2]` and get combined correctly, because
the accumulation is keyed by level, not by visit order.

The two arrays start empty and grow lazily: the first time the recursion reaches a new
depth, it appends a fresh slot to both `result` and `count` for that level; every later
visit to the same depth reuses the existing slot instead of appending again. Once the
whole tree is visited, a single final pass divides each level's sum by its count,
turning running totals into averages.

### The solution

![Example 1](images/1.jpg "Example1")

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

Walking it through `[3,9,20,null,null,15,7]` (expected `[3.0, 14.5, 11.0]`):

- Visit `3` (level 0): no slot exists yet, so slot 0 is created — `result=[3.0]`,
  `count=[1.0]`.

  ![Step 1: visit 3 (level 0), create slot 0, result[0]=3, count[0]=1](images/walkthrough-1.svg)

- Visit `9` (level 1): slot 1 doesn't exist yet either, so it's created —
  `result=[3.0, 9.0]`, `count=[1.0, 1.0]`.

  ![Step 2: visit 9 (level 1), create slot 1, result[1]=9, count[1]=1](images/walkthrough-2.svg)

- Visit `20` (level 1): slot 1 already exists from the `9` visit, so it's updated in
  place instead of appended: `result[1] += 20` → `29.0`, `count[1]++` → `2.0`. `20` is
  in a completely different branch than `9`, yet they both write into the same slot
  because they share a depth — that's the whole trick.

  ![Step 3: visit 20 (level 1), slot 1 exists, result[1] becomes 29, count[1] becomes 2](images/walkthrough-3.svg)

- Visit `15` (level 2): no slot 2 yet, so it's created — `result=[3.0, 29.0, 15.0]`,
  `count=[1.0, 2.0, 1.0]`.

  ![Step 4: visit 15 (level 2), create slot 2, result[2]=15, count[2]=1](images/walkthrough-4.svg)

- Visit `7` (level 2): slot 2 already exists, updated in place: `result[2] += 7` →
  `22.0`, `count[2]++` → `2.0`.

  ![Step 5: visit 7 (level 2), slot 2 exists, result[2] becomes 22, count[2] becomes 2](images/walkthrough-5.svg)

- Recursion done: `result=[3.0, 29.0, 22.0]`, `count=[1.0, 2.0, 2.0]`. The final loop
  divides each slot: `3.0/1.0 = 3.0`, `29.0/2.0 = 14.5`, `22.0/2.0 = 11.0` →
  `[3.0, 14.5, 11.0]` ✓.

  ![Step 6: divide result[i] by count[i] for every level, giving 3.0, 14.5, 11.0](images/walkthrough-6.svg)

**Complexity:** O(n) time — every node is visited exactly once, and the final division
pass is bounded by the number of levels. O(h) space for the recursion stack (h = tree
height), plus O(L) for the `result`/`count` slices, where L (number of levels) is at
most h + 1.

Full code: `easy_problems/601_700/average_of_levels_in_binary_tree/` in the repo.
