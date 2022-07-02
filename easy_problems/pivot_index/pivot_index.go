package pivot_index

func PivotIndex(nums []int) int {
	pivot, n := -1, len(nums)

	for i := 1; i < n; i++ {
		nums[i] += nums[i-1]
	}
	if nums[n-1]-nums[0] == 0 {
		return 0
	}

	for i := 1; i < n; i++ {
		left := nums[i-1]
		right := nums[n-1] - nums[i]

		if left == right {
			pivot = i
			break
		}
	}
	return pivot
}
