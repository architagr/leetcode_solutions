package uniquepaths

func uniquePaths(m int, n int) int {
	prev, curr := make([]int, n), make([]int, n)

	for i := 0; i < n; i++ {
		prev[i] = 1
	}
	for i := 1; i < m; i++ {
		curr[0] = 1
		for j := 1; j < n; j++ {
			curr[j] = prev[j] + curr[j-1]
		}
		copy(prev, curr)
	}
	return prev[n-1]
}
