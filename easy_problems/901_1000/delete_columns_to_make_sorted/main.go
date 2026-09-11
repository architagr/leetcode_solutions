package deletecolumnstomakesorted

func minDeletionSize(strs []string) int {
	n, m, count := len(strs), len(strs[0]), 0

	for j := 0; j < m; j++ {
		for i := 1; i < n; i++ {
			if strs[i][j] < strs[i-1][j] {
				count++
				break
			}
		}
	}
	return count
}
