package maximumsumofdistinctsubarrayswithlengthk

func maximumSubarraySum(nums []int, k int) int64 {
	k--
	m := make(map[int]int)
	sum, result := int64(0), int64(0)
	for i := 0; i < k; i++ {
		m[nums[i]]++
		sum += int64(nums[i])
	}
	for i := k; i < len(nums); i++ {
		m[nums[i]]++
		sum += int64(nums[i])
		if len(m) == k+1 {
			result = maxVal(result, sum)
		}
		m[nums[i-k]]--
		if m[nums[i-k]] == 0 {
			delete(m, nums[i-k])
		}
		sum -= int64(nums[i-k])
	}
	return result
}

func maxVal(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}
