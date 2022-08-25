package matrix_block_sum

import "math"

func MatrixBlockSum(mat [][]int, k int) [][]int {
	prefixSum := getPrefixSum(mat)
	ans := make([][]int, len(mat))
	n, m := len(mat), len(mat[0])

	for i := 0; i < n; i++ {
		r1, r2 := getParam(i, k, n)
		row := make([]int, m)
		for j := 0; j < m; j++ {
			c1, c2 := getParam(j, k, m)
			sum := prefixSum[r2][c2]

			if r1 > 0 {
				sum -= prefixSum[r1-1][c2]
			}
			if c1 > 0 {
				sum -= prefixSum[r2][c1-1]
			}
			if r1 > 0 && c1 > 0 {
				sum += prefixSum[r1-1][c1-1]
			}
			row[j] = sum
		}
		ans[i] = row
	}
	return ans
}

func getPrefixSum(mat [][]int) [][]int {
	ans := make([][]int, len(mat))

	for i := 0; i < len(mat); i++ {
		sum := 0
		ans[i] = make([]int, len(mat[0]))
		for j := 0; j < len(mat[0]); j++ {
			sum += mat[i][j]
			ans[i][j] = sum
		}
	}
	for j := 0; j < len(mat[0]); j++ {
		sum := 0
		for i := 0; i < len(mat); i++ {
			sum += ans[i][j]
			ans[i][j] = sum
		}
	}
	return ans
}

func getParam(i, k, max int) (int, int) {
	return int(math.Max(float64(0), float64(i-k))), int(math.Min(float64(i+k), float64(max-1)))
}
