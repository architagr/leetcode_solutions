package painthouse

func minCost(costs [][]int) int {
	n := len(costs)
	a, b, c := costs[n-1][0], costs[n-1][1], costs[n-1][2]
	for i := n - 2; i >= 0; i-- {
		costs[i][0] += minValue(b, c)
		costs[i][1] += minValue(a, c)
		costs[i][2] += minValue(a, b)
		a, b, c = costs[i][0], costs[i][1], costs[i][2]
	}

	return minValue(minValue(a, b), c)
}

func minValue(a, b int) int {
	if a < b {
		return a
	}
	return b
}
