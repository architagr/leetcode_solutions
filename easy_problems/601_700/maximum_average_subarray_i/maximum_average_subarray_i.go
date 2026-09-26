package maximumaveragesubarrayi

import "math"

func findMaxAverage(nums []int, k int) float64 {
	// From here k is an offset: nums[i-k] is the oldest element of the window
	// ending at i. The length comes back as k+1 in the final division.
	k--
	// Sums can be negative, so 0 would be a wrong starting best.
	max := math.MinInt
	sum := 0
	// Prime the window one element short, so the loop below can always do
	// add, check, remove in the same order, including on its first pass.
	for i := 0; i < k; i++ {
		sum += nums[i]
	}
	for i := k; i < len(nums); i++ {
		sum += nums[i]
		// Every window has the same length, so the largest sum is the largest
		// average; compare ints here and divide once at the end.
		if sum > max {
			max = sum
		}
		sum -= nums[i-k]
	}
	return float64(max) / float64(k+1)
}
