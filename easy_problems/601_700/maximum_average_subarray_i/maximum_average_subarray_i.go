package maximumaveragesubarrayi

import "math"

func findMaxAverage(nums []int, k int) float64 {
	k--
	max := math.MinInt
	sum := 0
	for i := 0; i < k; i++ {
		sum += nums[i]
	}
	for i := k; i < len(nums); i++ {
		sum += nums[i]
		if sum > max {
			max = sum
		}
		sum -= nums[i-k]
	}
	return float64(max) / float64(k+1)
}
