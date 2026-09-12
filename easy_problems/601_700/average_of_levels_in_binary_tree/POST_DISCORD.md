**365 Days of LeetCode Challenge — Day 27/365**
**Average of Levels in Binary Tree** (Easy)
🔗 https://leetcode.com/problems/average-of-levels-in-binary-tree/

Most solutions here drain the tree level by level with a BFS queue. I went with DFS instead, carrying the current depth as a parameter and using it as an index into two running-total arrays. Nodes from completely different branches still land in the same slot as long as they share a depth. Feels a little like cheating the first time you see it work.

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

Full walkthrough with step-by-step diagrams: https://github.com/architagr/leetcode_solutions/blob/main/easy_problems/601_700/average_of_levels_in_binary_tree/SOLUTION.md
