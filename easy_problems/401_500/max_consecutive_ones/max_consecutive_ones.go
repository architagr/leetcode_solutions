package max_consecutive_ones

func FindMaxConsecutiveOnes(nums []int) int {
	count, max, n := 0, 0, len(nums)
	for i := 0; i < n; i++ {
		if nums[i] == 0 {
			count = 0
		} else {
			count++
			if max < count {
				max = count
			}
		}
	}

	return max
}
