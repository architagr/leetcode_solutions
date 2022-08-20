package minimum_number_of_refueling_stops

import "math"

func MinRefuelStops(target, startFuel int, stations [][]int) int {
	n := len(stations)
	dp := make([]int64, n+1)
	dp[0] = int64(startFuel)
	for i := 0; i < n; i++ {
		for j := i + 1; j > 0; j-- {
			if dp[j-1] >= int64(stations[i][0]) {
				dp[j] = int64(math.Max(float64(dp[j]), float64(dp[j-1]+int64(stations[i][1]))))
			}
		}
	}
	for i := 0; i < n+1; i++ {
		if dp[i] >= int64(target) {
			return i
		}
	}
	return -1
}
