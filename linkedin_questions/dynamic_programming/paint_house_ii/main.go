package painthouseii

import "math"

func minCostII(costs [][]int) int {
	n := len(costs)
	for i := n - 2; i >= 0; i-- {
		for j := 0; j < len(costs[i]); j++ {
			x := math.MaxInt64
			for k := 0; k < len(costs[i]); k++ {
				if k != j {
					x = minValue(x, costs[i+1][k])
				}
			}
			costs[i][j] += x
		}
	}
	x := math.MaxInt64
	for k := 0; k < len(costs[0]); k++ {
		x = minValue(x, costs[0][k])
	}
	return x
}

func minValue(a, b int) int {
	if a < b {
		return a
	}
	return b
}
