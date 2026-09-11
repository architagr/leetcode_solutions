package minimumcosttoreacheveryposition

func minCosts(cost []int) []int {
	result := make([]int, len(cost))
	result[0] = cost[0]
	for i := 1; i < len(cost); i++ {
		if cost[i] < result[i-1] {
			result[i] = cost[i]
		} else {
			result[i] = result[i-1]
		}
	}
	return result
}
