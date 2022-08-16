package sort_array_by_increasing_frequency

import "sort"

func FrequencySort(nums []int) []int {
	m := make(map[int]int)

	for _, v := range nums {
		m[v]++
	}
	sort.Slice(nums, func(a, b int) bool {
		if m[nums[a]] == m[nums[b]] {
			return nums[a] > nums[b]
		}

		return m[nums[a]] < m[nums[b]]
	})

	return nums
}
