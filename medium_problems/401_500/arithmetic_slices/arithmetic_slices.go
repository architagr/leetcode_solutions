package arithmeticslices

func numberOfArithmeticSlices(nums []int) int {
	n := len(nums)
	if n < 3 {
		return 0
	}
	count, curr, currCount := 0, nums[1]-nums[0], 1

	for i := 2; i < n; i++ {
		if nums[i]-nums[i-1] == curr {
			currCount++
			continue
		}
		currCount--
		count += currCount * (currCount + 1) / 2
		currCount, curr = 1, nums[i]-nums[i-1]
	}
	currCount--
	count += currCount * (currCount + 1) / 2
	return count
}
