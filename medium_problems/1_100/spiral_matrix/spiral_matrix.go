package spiral_matrix

func SpiralOrder(matrix [][]int) []int {
	m, n, i, j, x := len(matrix), len(matrix[0]), 0, 0, 0
	result := make([]int, 0, m*n)
	for len(result)+1 < m*n {
		//row top
		for ; j < n-1-x; j++ {
			result = append(result, matrix[i][j])
		}

		//right most column
		for ; i < m-1-x; i++ {
			result = append(result, matrix[i][j])
		}

		// last row
		for ; len(result) < m*n && j > x-0; j-- {
			result = append(result, matrix[i][j])
		}
		// left column
		for ; len(result) < m*n && i > x-0; i-- {
			result = append(result, matrix[i][j])
		}
		x++
		i, j = x, x
	}
	if len(result) < m*n {
		result = append(result, matrix[i][j])
	}

	return result
}
