package top_k_frequent_elements

import "sort"

func TopKFrequent(nums []int, k int) []int {
	m := make(map[int]int)

	for _, v := range nums {
		m[v]++
	}
	type frequency struct {
		count, ele int
	}
	arr := make([]frequency, len(m))
	j := 0
	for key, value := range m {
		arr[j] = frequency{
			count: value,
			ele:   key,
		}
		j++
	}
	sort.Slice(arr, func(i, j int) bool { return arr[i].count > arr[j].count })
	ans := make([]int, 0, k)
	for i := 0; i < len(arr) && k > 0; i, k = i+1, k-1 {
		ans = append(ans, arr[i].ele)
	}
	return ans
}
