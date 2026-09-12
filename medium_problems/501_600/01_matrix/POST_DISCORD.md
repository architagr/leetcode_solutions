**365 Days of LeetCode Challenge — Day 41/365**
**01 Matrix** (Medium)
🔗 https://leetcode.com/problems/01-matrix/

This is yesterday's problem. Rotting Oranges: how long until an orange rots, every rotten one spreading at once. Today: how far is a cell from the nearest zero, every zero a source. Same question, and multi-source BFS solves it directly.

This solution uses no queue at all, which is why it is worth reading.

The definition is circular - a cell's distance is 1 + the smallest among its four neighbours, but every neighbour depends on the cell. No order to evaluate it in.

Direction breaks the circle. Any shortest path leaves in one of four directions; split them into up-or-left and down-or-right.

Sweep top-left to bottom-right: everything above and left is already final, so you can compute the best distance for paths starting up or left.

```go
minNeighbor := m * n
if i > 0 { minNeighbor = minValue(minNeighbor, result[i-1][j]) }
if j > 0 { minNeighbor = minValue(minNeighbor, result[i][j-1]) }
result[i][j] = minNeighbor + 1
```

Then sweep bottom-right to top-left for the other two, and here is the line that matters:

```go
result[i][j] = minValue(result[i][j], minNeighbor+1)
```

min with the existing value, NOT an assignment. Overwrite and you discard everything pass one got right.

O(m*n) time, no queue.

Full walkthrough: https://github.com/architagr/leetcode_solutions/blob/main/medium_problems/501_600/01_matrix/SOLUTION.md
