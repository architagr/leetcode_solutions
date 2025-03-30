package besttimetobuyandsellstockiv

func maxProfit(k int, prices []int) int {
	dpTable := make([][][]int, len(prices))
	for i := 0; i < len(prices); i++ {
		dpTable[i] = make([][]int, k+1)
		for j := 0; j <= k; j++ {
			dpTable[i][j] = make([]int, 2)
		}
	}

	return dp(0, k, 0, &dpTable, prices)
}

func dp(i, remainingDays, holding int, dpTable *[][][]int, prices []int) int {
	if i == len(prices) || remainingDays == 0 {
		return 0
	}

	if (*dpTable)[i][remainingDays][holding] == 0 {
		doNothing := dp(i+1, remainingDays, holding, dpTable, prices)
		doSomething := 0

		if holding == 1 {
			// sell
			doSomething = prices[i] + dp(i+1, remainingDays-1, 0, dpTable, prices)
		} else {
			// buy
			doSomething = -prices[i] + dp(i+1, remainingDays, 1, dpTable, prices)
		}
		(*dpTable)[i][remainingDays][holding] = maxValue(doNothing, doSomething)
	}

	return (*dpTable)[i][remainingDays][holding]
}

func maxValue(a, b int) int {
	if a > b {
		return a
	}
	return b
}
