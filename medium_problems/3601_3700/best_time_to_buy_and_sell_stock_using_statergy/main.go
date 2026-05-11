package besttimetobuyandsellstockusingstatergy

func maxProfit(prices []int, strategy []int, k int) int64 {
	original := make([]int64, len(prices))
	original[0] = int64(prices[0]) * int64(strategy[0])
	for i := 1; i < len(prices); i++ {
		original[i] = (int64(prices[i]) * int64(strategy[i])) + original[i-1]
	}
	prices = prefixSum(prices)
	ans := original[len(prices)-1]

	for left, right := 0, k-1; right < len(prices); right, left = right+1, left+1 {
		a := int64(0)
		if left > 0 {
			a += original[left-1]
		}
		b := int64(prices[right]) - int64(prices[(right+left)/2])
		c := original[len(prices)-1] - original[right]

		ans = max(ans, a+b+c)
	}

	return ans
}

func max(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}
func prefixSum(nums []int) []int {
	result := make([]int, len(nums))

	result[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		result[i] = result[i-1] + nums[i]
	}
	return result
}
