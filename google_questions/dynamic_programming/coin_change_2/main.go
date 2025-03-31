package coinchange2

func change(amount int, coins []int) int {
	dpTable := make([][]int, len(coins))
	for i := 0; i < len(coins); i++ {
		dpTable[i] = make([]int, amount+1)
		for j := 0; j < amount+1; j++ {
			dpTable[i][j] = -1
		}
	}
	return ways(0, amount, &dpTable, coins)
}

func ways(i, amount int, dpTable *[][]int, coins []int) int {
	if amount == 0 {
		return 1
	}
	if i == len(coins) {
		return 0
	}
	if (*dpTable)[i][amount] == -1 {
		if coins[i] > amount {
			(*dpTable)[i][amount] = ways(i+1, amount, dpTable, coins)
		} else {
			(*dpTable)[i][amount] = ways(i, amount-coins[i], dpTable, coins) + ways(i+1, amount, dpTable, coins)
		}
	}

	return (*dpTable)[i][amount]
}
