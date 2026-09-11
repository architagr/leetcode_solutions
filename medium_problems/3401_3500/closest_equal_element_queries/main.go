package closestequalelementqueries

func solveQueries(nums []int, queries []int) []int {
	arrayLen := len(nums)
	indexMap := buildIndexMap(nums)

	response := make([]int, len(queries))
	for i, qIndex := range queries {
		n := nums[qIndex]

		indices := indexMap[n]
		if len(indices) == 1 {
			response[i] = -1
			continue
		}
		countIndices := len(indices)
		l := -1
		for ii, ll := range indices {
			if ll == qIndex {
				l = ii
				break
			}
		}

		next := absValue(indices[((countIndices+l+1)%countIndices)] - indices[l])
		prev := absValue(indices[((countIndices+l-1)%countIndices)] - indices[l])
		nextMin := minValue(next, arrayLen-next)
		prevMin := minValue(prev, arrayLen-prev)
		response[i] = minValue(nextMin, prevMin)
	}
	return response
}

func absValue(a int) int {
	if a < 0 {
		a *= -1
	}
	return a
}
func minValue(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func buildIndexMap(nums []int) map[int][]int {
	response := make(map[int][]int, len(nums))
	for i, n := range nums {
		response[n] = append(response[n], i)
	}
	return response
}
