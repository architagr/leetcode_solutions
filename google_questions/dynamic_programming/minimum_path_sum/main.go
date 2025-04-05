package minimumpathsum

func minPathSum(grid [][]int) int {

	n, m := len(grid), len(grid[0])
	for i := 1; i < m; i++ {
		grid[0][i] += grid[0][i-1]
	}
	for i := 1; i < n; i++ {
		grid[i][0] += grid[i-1][0]
	}
	for i := 1; i < n; i++ {
		for j := 1; j < m; j++ {
			grid[i][j] += minValue(grid[i-1][j], grid[i][j-1])
		}
	}
	return grid[n-1][m-1]
}

func minValue(a, b int) int {
	if a < b {
		return a
	}
	return b
}
