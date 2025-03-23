package cheapestflightswithinkstops

import "math"

func findCheapestPrice(n int, flights [][]int, src int, dst int, k int) int {

	if src == dst {
		return 0
	}
	prev, current := make([]int, n), make([]int, n)
	for i := 0; i < n; i++ {
		prev[i] = math.MaxInt
		current[i] = math.MaxInt
	}
	prev[src] = 0
	for i := 0; i < k+1; i++ {
		current[src] = 0
		for _, f := range flights {
			src, dest, cost := f[0], f[1], f[2]
			if prev[src] < math.MaxInt {
				current[dest] = minValue(current[dest], prev[src]+cost)
			}
		}
		prev = append([]int{}, current...)
	}

	if current[dst] == math.MaxInt {
		return -1
	}
	return current[dst]
}

func minValue(a, b int) int {
	if a < b {
		return a
	}
	return b
}
