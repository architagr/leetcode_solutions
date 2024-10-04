package leftandrightsumdifferences

func leftRightDifference(nums []int) []int {

	result := make([]int, len(nums))
	sum := 0
	for i := range nums {
		result[i] = sum

		sum += nums[i]
	}

	sum = 0
	for i := len(nums) - 1; i >= 0; i-- {
		result[i] = abs(result[i] - sum)

		sum += nums[i]
	}

	return result
}

func abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}
