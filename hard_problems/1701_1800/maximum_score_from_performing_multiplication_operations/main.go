package maximumscorefromperformingmultiplicationoperations

func maximumScore(nums []int, multipliers []int) int {
	m := len(multipliers)
	dpTable := make([][]int, m*m)
	for i := 0; i < m; i++ {
		dpTable[i] = make([]int, m)
	}
	return dp(0, 0, nums, multipliers, &dpTable)
}

func dp(i, left int, nums, multipliers []int, dpTable *[][]int) int {
	if i == len(multipliers) {
		return 0
	}

	mult := multipliers[i]
	right := len(nums) - 1 - (i - left)

	if (*dpTable)[i][left] == 0 {
		// Recurrence relation
		(*dpTable)[i][left] = maxValue(mult*nums[left]+dp(i+1, left+1, nums, multipliers, dpTable),
			mult*nums[right]+dp(i+1, left, nums, multipliers, dpTable))
	}

	return (*dpTable)[i][left]
}

func maxValue(a, b int) int {
	if a > b {
		return a
	}
	return b
}
