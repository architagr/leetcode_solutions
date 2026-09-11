package dungeongame

import "math"

var dp [][]int

func calculateMinimumHP(A [][]int) int {
	n, m := len(A), len(A[0])
	dp = make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, m)
	}
	return calc(A, 0, 0)
}

func calc(A [][]int, i, j int) int {
	n, m := len(A), len(A[0])
	if i == n-1 && j == m-1 {
		return maxValue(1, (1 - A[i][j]))
	}
	if i >= n || j >= m {
		return math.MaxInt
	}
	if dp[i][j] == 0 {
		top := calc(A, i+1, j)
		left := calc(A, i, j+1)
		dp[i][j] = maxValue(1, minValue(top, left)-A[i][j])
	}
	return dp[i][j]
}
func minValue(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func maxValue(a, b int) int {
	if a > b {
		return a
	}
	return b
}
