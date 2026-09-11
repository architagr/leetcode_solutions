package min_max_game

func MinMaxGame(nums []int) int {
	n := len(nums)
	if n == 1 {
		return nums[0]
	}
	result := make([]int, n/2)

	for i := 0; i < n/2; i++ {

		if i%2 == 0 {
			result[i] = minValue(nums[2*i], nums[(2*i)+1])
		} else {
			result[i] = maxValue(nums[2*i], nums[(2*i)+1])
		}
	}
	return MinMaxGame(result)
}

func maxValue(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func minValue(a, b int) int {
	if a < b {
		return a
	}
	return b
}
