package minimum_average_difference

import "math"

func minimumAverageDifference(nums []int) int {
	total := 0

	for i := 0; i < len(nums); i++ {
		total += nums[i]
	}

	sum, min, index := 0, math.MaxInt, -1
	i := 0
	for ; i < len(nums)-1; i++ {
		sum += nums[i]
		x := (sum / (i + 1))
		x -= ((total - sum) / (len(nums) - i - 1))
		x = abs(x)
		if x < min {
			min = x
			index = i
		}

	}
	sum += nums[i]
	x := (sum / (i + 1))
	x = abs(x)
	if x < min {
		index = i
	}
	return index
}

func abs(a int) int {
	if a < 0 {
		a *= -1
	}
	return a
}
