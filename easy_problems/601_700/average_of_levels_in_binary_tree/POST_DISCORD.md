**365 Days of LeetCode Challenge — Day 16/365**
**Average of Levels in Binary Tree** (Easy)
🔗 https://leetcode.com/problems/average-of-levels-in-binary-tree/

**Intuition:** Most solutions here use BFS, draining the tree one level at a time with
a queue. I went with DFS instead: it carries the current depth as a parameter and uses
that as an index into two running-total arrays (`result[level]`, `count[level]`), so
nodes from completely different branches still accumulate into the same slot as long
as they're at the same depth. Feels a little like cheating the first time you see it
work.

![Example 1](images/1.jpg "Example1")

**Full solution:**
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

**Walkthrough** on `[3,9,20,null,null,15,7]` (expected `[3.0, 14.5, 11.0]`):

- Visit `3` (level 0): slot 0 created → `result=[3.0]`, `count=[1.0]`

![Step 1: visit 3 (level 0), create slot 0, result[0]=3, count[0]=1](images/walkthrough-1.svg)

- Visit `9` (level 1): slot 1 created → `result=[3.0, 9.0]`, `count=[1.0, 1.0]`

![Step 2: visit 9 (level 1), create slot 1, result[1]=9, count[1]=1](images/walkthrough-2.svg)

- Visit `20` (level 1): slot 1 already exists (from `9`), so it's updated in place:
  `result[1] += 20` → `29.0`, `count[1]++` → `2.0`

![Step 3: visit 20 (level 1), slot 1 exists, result[1] becomes 29, count[1] becomes 2](images/walkthrough-3.svg)

- Visit `15` (level 2): slot 2 created → `result=[3.0, 29.0, 15.0]`,
  `count=[1.0, 2.0, 1.0]`

![Step 4: visit 15 (level 2), create slot 2, result[2]=15, count[2]=1](images/walkthrough-4.svg)

- Visit `7` (level 2): slot 2 already exists, updated in place: `result[2] += 7` →
  `22.0`, `count[2]++` → `2.0`

![Step 5: visit 7 (level 2), slot 2 exists, result[2] becomes 22, count[2] becomes 2](images/walkthrough-5.svg)

- Final pass divides each slot by its count: `3.0/1 = 3.0`, `29.0/2 = 14.5`,
  `22.0/2 = 11.0`, giving `[3.0, 14.5, 11.0]`

![Step 6: divide result[i] by count[i] for every level, giving 3.0, 14.5, 11.0](images/walkthrough-6.svg)

O(n) time, every node gets visited once. O(h) space for the recursion stack, plus O(L)
for the level arrays (L is the number of levels, at most h + 1).
