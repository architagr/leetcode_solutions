package min_cost_climbing_stairs

func MinCostClimbingStairs(cost []int) int {
	n := len(cost)
	for i := 2; i < n; i++ {
		if cost[i-2] < cost[i-1] {
			cost[i] = cost[i] + cost[i-2]
		} else {
			cost[i] = cost[i] + cost[i-1]
		}
	}
	if cost[n-2] < cost[n-1] {
		return cost[n-2]
	}
	return cost[n-1]

}
func MinCostClimbingStairs1(cost []int) int {
	cost = append(cost, 0)
	x := findMinCost(0, cost)
	y := findMinCost(1, cost)

	if x < y {
		return x
	}
	return y
}

func findMinCost(start int, cost []int) int {
	n, k := len(cost), 2
	nums := make([]int, n)
	nums[start] = cost[start]
	queue := []int{start}

	for i := queue[0] + 1; i < n; i++ {
		if queue[0] < i-k {
			queue = queue[1:]
		}
		nums[i] = cost[i] + nums[queue[0]]
		for len(queue) > 0 && nums[queue[len(queue)-1]] > nums[i] {
			queue = queue[:len(queue)-1]
		}

		queue = append(queue, i)
	}

	return nums[n-1]
}
