package find_closest_number_to_zero

import "math"

func FindClosestNumber(nums []int) int {

	min, ans := math.MaxInt32, 0

	for i := 0; i < len(nums); i++ {
		num := nums[i]
		if num < 0 {
			num *= -1
		}
		if num < min {
			ans = nums[i]
			min = num
			continue
		}
		if num == min {
			if nums[i] > ans {
				ans = nums[i]
			}
		}
	}
	return ans
}
