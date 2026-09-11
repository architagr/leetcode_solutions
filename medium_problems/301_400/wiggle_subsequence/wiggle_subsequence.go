package wiggle_subsequence

func WiggleMaxLength(nums []int) int {
	lastDiff, n, currDiff := 0, len(nums), 0
	count := 1
	if n == 1 {
		return 1
	}

	lastDiff = nums[1] - nums[0]
	if lastDiff != 0 {
		count = 2
	}
	for i := 2; i < n; i++ {
		currDiff = nums[i] - nums[i-1]

		if (currDiff > 0 && lastDiff <= 0) || (currDiff < 0 && lastDiff >= 0) {
			count++
			lastDiff = currDiff
		}
	}
	return count
}
