package besttimetobuyandsellstock

func maxProfit(prices []int) int {
	profit := 0
	min := prices[0]
	for i := 1; i < len(prices); i++ {
		if min > prices[i] {
			min = prices[i]
		} else {
			profit = maxVal(profit, prices[i]-min)
		}
	}
	return profit
}

func maxVal(a, b int) int {
	if a > b {
		return a
	}
	return b
}
