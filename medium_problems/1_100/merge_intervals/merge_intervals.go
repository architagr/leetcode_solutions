package merge_intervals

import (
	"sort"
)

func Merge(intervals [][]int) [][]int {
	result := make([][]int, 0)
	n, k := len(intervals), 1
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] == intervals[j][0] {
			return intervals[i][1] < intervals[j][1]
		}
		return intervals[i][0] < intervals[j][0]
	})
	result = append(result, intervals[0])
	for i := 1; i < n; i++ {
		if intervals[i][0] <= result[k-1][1] {
			result[k-1][1] = maxValue(intervals[i][1], result[k-1][1])
			continue
		}
		result = append(result, intervals[i])
		k++

	}
	return result
}

func maxValue(a, b int) int {
	if a > b {
		return a
	}
	return b
}
