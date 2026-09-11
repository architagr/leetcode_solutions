package merge_similar_items

import "sort"

func MergeSimilarItems(items1 [][]int, items2 [][]int) [][]int {
	mapItems := make(map[int]int)

	for i := 0; i < len(items1); i++ {
		mapItems[items1[i][0]] = items1[i][1]
	}
	for i := 0; i < len(items2); i++ {
		val, ok := mapItems[items2[i][0]]

		if ok {
			val += items2[i][1]
		} else {
			val = items2[i][1]
		}
		mapItems[items2[i][0]] = val
	}
	result := make([][]int, len(mapItems))
	k := 0
	for key, val := range mapItems {
		result[k] = []int{key, val}
		k++
	}
	sort.Slice(result, func(i, j int) bool {
		return result[i][0] < result[j][0]
	})
	return result
}
