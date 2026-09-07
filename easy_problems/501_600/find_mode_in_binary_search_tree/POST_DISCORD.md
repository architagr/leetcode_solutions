**365 Days of LeetCode Challenge — Day 29/365**
**Find Mode in Binary Search Tree** (Easy)
🔗 https://leetcode.com/problems/find-mode-in-binary-search-tree/

The "BST" part is kind of beside the point here. This approach ignores the ordering completely, tallies how many times every value shows up in a hashmap, then reads off whichever value or values hit the highest count. Ties get handled for free, since the second pass grabs everything matching the max.

```go
func findMode(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	data := make(map[int]int)
	data = getCnt(root, data)

	res := make([]int, 0)
	cnt := 0
	for _, v := range data {
		if v > cnt {
			cnt = v
		}
	}

	for k, v := range data {
		if v == cnt {
			res = append(res, k)
		}
	}
	return res
}

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

O(n) time, O(n) space for the map plus O(h) for the recursion stack.

Full walkthrough with step-by-step diagrams: https://github.com/architagr/leetcode_solutions/blob/main/easy_problems/501_600/find_mode_in_binary_search_tree/SOLUTION.md
