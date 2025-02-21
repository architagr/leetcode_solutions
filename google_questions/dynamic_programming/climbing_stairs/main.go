package climbingstairs

func climbStairs(n int) int {
	dp := make([]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = -1
	}
	dp[0] = 1
	dp[1] = 1
	febbonic(n, &dp)
	return dp[n]
}

func febbonic(n int, dp *[]int) {
	if (*dp)[n] >= 0 {
		return
	}
	febbonic(n-1, dp)
	febbonic(n-2, dp)
	(*dp)[n] = (*dp)[n-1] + (*dp)[n-2]
}
