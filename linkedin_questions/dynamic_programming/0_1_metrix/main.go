package metrix

func updateMatrix(mat [][]int) [][]int {
	result := make([][]int, len(mat))
	n := len(mat)
	m := len(mat[0])
	for i := 0; i < n; i++ {
		result[i] = make([]int, m)
		for j := 0; j < m; j++ {
			result[i][j] = mat[i][j]
		}
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if result[i][j] == 0 {
				continue
			}
			minNeighbor := m * n
			if i > 0 {
				minNeighbor = minValue(minNeighbor, result[i-1][j])
			}

			if j > 0 {
				minNeighbor = minValue(minNeighbor, result[i][j-1])
			}

			result[i][j] = minNeighbor + 1
		}
	}
	for i := n - 1; i >= 0; i-- {
		for j := m - 1; j >= 0; j-- {
			if result[i][j] == 0 {
				continue
			}

			minNeighbor := m * n
			if i < n-1 {
				minNeighbor = minValue(minNeighbor, result[i+1][j])
			}

			if j < m-1 {
				minNeighbor = minValue(minNeighbor, result[i][j+1])
			}

			result[i][j] = minValue(result[i][j], minNeighbor+1)
		}
	}
	return result
}

func minValue(a, b int) int {
	if a < b {
		return a
	}
	return b
}
