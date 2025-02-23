package partitionlabels

type index struct {
	start, end int
}

func partitionLabels(s string) []int {
	result := make([]int, 0)
	indexMap := make([]index, 0, 26)
	letterMap := make(map[byte]int, 26)
	for i := 0; i < len(s); i++ {
		idx, ok := letterMap[s[i]]
		if !ok {
			indexMap = append(indexMap, index{
				start: i,
				end:   i,
			})
			letterMap[s[i]] = len(indexMap) - 1
		} else {
			indexMap[idx].end = i
		}
	}

	currStart, currEnd := indexMap[0].start, indexMap[0].end
	for i := 1; i < len(indexMap); i++ {
		if indexMap[i].start < currEnd {
			currStart = minVal(currStart, indexMap[i].start)
			currEnd = maxVal(currEnd, indexMap[i].end)
		} else {
			result = append(result, currEnd-currStart+1)
			currStart, currEnd = indexMap[i].start, indexMap[i].end
		}
	}
	result = append(result, currEnd-currStart+1)
	return result
}

func minVal(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func maxVal(a, b int) int {
	if a > b {
		return a
	}
	return b
}
