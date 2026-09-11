package findthescoreofallprefixesofanarray

func findPrefixScore(nums []int) []int64 {
	max := nums[0]
	var val int64
	l := len(nums)
	result := make([]int64, l)
	for i := 0; i < l; i++ {
		if nums[i] > max {
			max = nums[i]
		}
		val = int64(nums[i]) + int64(max) + val
		result[i] = val
	}
	return result
}
