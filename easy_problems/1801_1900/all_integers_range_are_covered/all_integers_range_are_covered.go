package allintegersrangearecovered

import "sort"

func isCovered(ranges [][]int, left int, right int) bool {
	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i][0] < ranges[j][0]
	})
	isRightPresent := false

	for _, rangeData := range ranges {
		if left >= rangeData[0] && left <= rangeData[1] {
			left = minValue(rangeData[1]+1, right)
		}
		if right >= rangeData[0] && right <= rangeData[1] {
			isRightPresent = true
		}

	}
	return isRightPresent && left == right
}
func minValue(a, b int) int {
	if a < b {
		return a
	}
	return b
}
