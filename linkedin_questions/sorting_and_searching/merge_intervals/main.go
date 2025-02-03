package mergeintervals

import "sort"

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] <= intervals[j][0]
	})
	result := make([][]int, 1, len(intervals))
	result[0] = intervals[0]
	for i, k := 1, 0; i < len(intervals); i, k = i+1, k+1 {
		for i < len(intervals) && result[k][1] >= intervals[i][0] {
			if result[k][1] <= intervals[i][1] {
				result[k][1] = intervals[i][1]
			}
			i++
		}
		if i < len(intervals) {
			result = append(result, intervals[i])
		}
	}
	return result
}
