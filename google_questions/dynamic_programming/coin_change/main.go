package coinchange

func coinChange(coins []int, amount int) int {
	max := amount + 1
	dp := make([]int, max)
	for i := 0; i < max; i++ {
		dp[i] = max
	}
	dp[0] = 0
	for i := 1; i <= amount; i++ {
		for j := 0; j < len(coins); j++ {
			if coins[j] <= i {
				dp[i] = minValue(dp[i], dp[i-coins[j]]+1)
			}
		}
	}
	if dp[amount] > amount {
		return -1
	}
	return dp[amount]
}

func minValue(a, b int) int {
	if a < b {
		return a
	}
	return b
}
