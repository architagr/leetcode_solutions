package maximumproductsubarray

func maxProduct(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	result, max, min := nums[0], nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		temp := maxVal(nums[i], maxVal(max*nums[i], min*nums[i]))
		min = minVal(nums[i], minVal(max*nums[i], min*nums[i]))
		max = temp

		result = maxVal(max, result)
	}
	return result
}
func minVal(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func maxVal(a, b int) int {
	if a > b {
		return a
	}
	return b
}
