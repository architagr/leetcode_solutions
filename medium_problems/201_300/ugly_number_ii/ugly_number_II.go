package ugly_number_II

func NthUglyNumber(n int) int {
	dp := make([]int, n+1)
	count_2, count_3, count_5 := 1, 1, 1

	min_val, val_2, val_3, val_5 := 0, 0, 0, 0

	dp[1] = 1 //ist ugly number is 1
	for i := 2; i <= n; i++ {
		val_2 = 2 * dp[count_2]
		val_3 = 3 * dp[count_3]
		val_5 = 5 * dp[count_5]
		min_val = min(val_2, min(val_3, val_5))
		if min_val == val_2 {
			count_2++
		}
		if min_val == val_3 {
			count_3++
		}
		if min_val == val_5 {
			count_5++
		}
		dp[i] = min_val
	}
	return dp[n]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
