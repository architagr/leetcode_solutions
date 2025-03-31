package besttimetobuyandsellstockwithcooldown

func maxProfit(prices []int) int {
	dpTable := make([][]int, len(prices))
	for i := 0; i < len(prices); i++ {
		dpTable[i] = make([]int, 2)

	}

	return dp(0, 0, &dpTable, prices)
}

func dp(i, holding int, dpTable *[][]int, prices []int) int {
	if i >= len(prices) {
		return 0
	}

	if (*dpTable)[i][holding] == 0 {
		doNothing := dp(i+1, holding, dpTable, prices)
		doSomething := 0

		if holding == 1 {
			// sell
			doSomething = prices[i] + dp(i+2, 0, dpTable, prices)
		} else {
			// buy
			doSomething = -prices[i] + dp(i+1, 1, dpTable, prices)
		}
		(*dpTable)[i][holding] = maxValue(doNothing, doSomething)
	}

	return (*dpTable)[i][holding]
}

func maxValue(a, b int) int {
	if a > b {
		return a
	}
	return b
}
