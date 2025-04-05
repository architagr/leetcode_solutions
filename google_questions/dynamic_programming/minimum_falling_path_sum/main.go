package minimumfallingpathsum

import "math"

func minFallingPathSum(matrix [][]int) int {
	n, m := len(matrix), len(matrix[0])
	for i := 1; i < n; i++ {
		for j := 0; j < m; j++ {
			left, center, right := math.MaxInt, matrix[i-1][j], math.MaxInt
			if j > 0 {
				left = matrix[i-1][j-1]
			}
			if j < m-1 {
				right = matrix[i-1][j+1]
			}
			matrix[i][j] += minValue(minValue(left, center), right)
		}
	}
	min := math.MaxInt

	for i := 0; i < m; i++ {
		min = minValue(min, matrix[n-1][i])
	}
	return min
}

func minValue(a, b int) int {
	if a < b {
		return a
	}
	return b
}
