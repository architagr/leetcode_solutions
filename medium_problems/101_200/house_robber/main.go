package houserobber

func rob(nums []int) int {
	return dp(len(nums)-1, nums, map[int]int{})
}

func dp(i int, nums []int, dpState map[int]int) int {
	if i == 0 {
		dpState[i] = nums[i]
	} else if i == 1 {
		dpState[i] = maxValue(nums[i], nums[i-1])
	} else if _, found := dpState[i]; !found {
		dpState[i] = maxValue(dp(i-1, nums, dpState), dp(i-2, nums, dpState)+nums[i])
	}

	return dpState[i]
}

func maxValue(a, b int) int {
	if a > b {
		return a
	}
	return b
}
