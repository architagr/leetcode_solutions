**365 Days of LeetCode Challenge — Day 39/365**
**Number of Provinces** (Medium)
🔗 https://leetcode.com/problems/number-of-provinces/

The counting is day 37's, unchanged. What changes is the input format, and this is the third graph representation in four days: day 36 an edge list, day 37 a grid with edges implied by position, today a matrix. All three become the same adjacency list.

The trap: this matrix is NOT a grid. In day 37 cell (i,j) was a place. Here cell (i,j) is a relationship between city i and city j - an n x n matrix describing n nodes, not n squared. You traverse cities and consult the matrix, not the reverse.

```go
for i := 0; i < len(isConnected); i++ {
	for j := i + 1; j < len(isConnected[i]); j++ {
		if isConnected[i][j] == 1 {
			ajList[i] = append(ajList[i], j)
			ajList[j] = append(ajList[j], i)
		}
	}
}
```

`j := i + 1` reads only above the diagonal. The matrix is symmetric so every relationship is stored twice, and starting at i+1 also skips the diagonal where a city is "connected" to itself.

One contrast worth noticing: day 36 marked nodes visited on ENQUEUE, this marks on DEQUEUE. Both correct. The dequeue version lets a node be pushed several times, which never shows in the output - only in how much work the queue does.

O(n^2) time, dominated by reading the matrix.

Full walkthrough: https://github.com/architagr/leetcode_solutions/blob/main/medium_problems/501_600/number_of_provinces/SOLUTION.md
