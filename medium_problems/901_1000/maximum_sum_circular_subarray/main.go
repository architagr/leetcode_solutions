package maximumsumcircularsubarray

func maxSubarraySumCircular(nums []int) int {
	curMax, curMin, totalSum := 0, 0, 0

	maxSum, minSum := nums[0], nums[0]

	for _, num := range nums {
		// Normal Kadane's
		curMax = maxValue(curMax, 0) + num
		maxSum = maxValue(maxSum, curMax)

		// Kadane's but with min to find minimum subarray
		curMin = minValue(curMin, 0) + num
		minSum = minValue(minSum, curMin)

		totalSum += num
	}

	if totalSum == minSum {
		return maxSum
	}

	return maxValue(maxSum, totalSum-minSum)
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
