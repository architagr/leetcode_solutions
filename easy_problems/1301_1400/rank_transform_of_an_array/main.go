package ranktransformofanarray

import "sort"

func arrayRankTransform(arr []int) []int {
	n := len(arr)
	temp := make([]int, 0, n)
	temp = append(temp, arr...)
	sort.Ints(temp)
	m := make(map[int]int)
	rank := 1
	for _, val := range temp {
		if _, ok := m[val]; ok {
			continue
		}
		m[val] = rank
		rank++
	}
	result := make([]int, n)
	for i, val := range arr {
		result[i] = m[val]
	}
	return result
}
