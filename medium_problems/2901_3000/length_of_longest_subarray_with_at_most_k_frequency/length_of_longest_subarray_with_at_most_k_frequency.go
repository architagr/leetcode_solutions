package lengthoflongestsubarraywithatmostkfrequency

func maxSubarrayLength(nums []int, k int) int {
	max := 0
	start := 0
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		m[nums[i]]++
		if m[nums[i]] <= k {
			max = maxVal(max, i-start+1)
			continue
		}

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
