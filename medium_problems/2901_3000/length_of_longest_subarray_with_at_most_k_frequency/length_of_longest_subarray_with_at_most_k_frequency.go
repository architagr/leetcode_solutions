package lengthoflongestsubarraywithatmostkfrequency

func maxSubarrayLength(nums []int, k int) int {
	max := 0
	start := 0
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		m[nums[i]]++
		// Only the value that just arrived can have gone over k.
		if m[nums[i]] <= k {
			max = maxVal(max, i-start+1)
			continue
		}

		// No record after shrinking: [start, i-1] was valid and already
		// recorded, and the shrunk window starts later, so it's no longer.
		for ; start < len(nums) && m[nums[i]] > k && start <= i; start++ {
			m[nums[start]]--
		}
	}

	return max
}

func maxVal(a, b int) int {
	if a > b {
		return a
	}
	return b
}
