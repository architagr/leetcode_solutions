package maximumerasurevalue

func maximumUniqueSubarray(nums []int) int {
	sum, max, left := 0, 0, 0
	m := make(map[int]int, len(nums))
	for right, val := range nums {
		if _, ok := m[val]; ok {
			max = maxVal(max, sum)
			for ; left <= m[val]; left++ {
				delete(m, nums[left])
				sum -= nums[left]
			}
		}
		m[val] = right
		sum += val
	}
	return maxVal(max, sum)
}

func maxVal(a, b int) int {
	if a > b {
		return a
	}
	return b
}
