package numberofwaystosplitarray

func waysToSplitArray(nums []int) int {
	n, count, sum := len(nums), 0, 0
	prefixSum := make([]int, n)

	for i := 0; i < n; i++ {
		sum += nums[i]
		prefixSum[i] = sum
	}

	for i := 0; i < n-1; i++ {
		if 2*prefixSum[i] >= sum {
			count++
		}
	}
	return count
}
