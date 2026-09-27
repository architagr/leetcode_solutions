package maximumerasurevalue

func maximumUniqueSubarray(nums []int) int {
	sum, max, left := 0, 0, 0
	m := make(map[int]int, len(nums))
	for right, val := range nums {
		if _, ok := m[val]; ok {
			// All values are positive, so the sum only grew since the last
			// repeat; record it before the window shrinks.
			max = maxVal(max, sum)
			// Remove everything up to and including the earlier copy of val.
			for ; left <= m[val]; left++ {
				delete(m, nums[left])
				sum -= nums[left]
			}
		}
		m[val] = right
		sum += val
	}
	// The final window never met a repeat, so record it here.
	return maxVal(max, sum)
}

func maxVal(a, b int) int {
	if a > b {
		return a
	}
	return b
}
