package max_area_cake_piece

import (
	"sort"
)

func MaxArea(h int, w int, horizontalCuts []int, verticalCuts []int) int {
	sort.Ints(horizontalCuts)
	sort.Ints(verticalCuts)
	horizontalCuts = append(horizontalCuts, h)
	verticalCuts = append(verticalCuts, w)
	maxH, maxV, n, m := horizontalCuts[0], verticalCuts[0], len(horizontalCuts), len(verticalCuts)

	for i := 1; i < n; i++ {
		maxH = maxValue(maxH, horizontalCuts[i]-horizontalCuts[i-1])
	}

	for i := 1; i < m; i++ {
		maxV = maxValue(maxV, verticalCuts[i]-verticalCuts[i-1])
	}
	return (maxH * maxV) % 1000000007
}

func maxValue(a, b int) int {
	if a > b {
		return a
	}
	return b
}
