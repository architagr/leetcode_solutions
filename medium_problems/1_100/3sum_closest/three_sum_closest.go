package three_sum_closest

import (
	"math"
	"sort"
)

func ThreeSumClosest(nums []int, target int) int {
	min, sum := math.MaxInt32, 0
	sort.Ints(nums)
	n := len(nums)

	for i := 0; i < n; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		l, r := i+1, n-1
		for l < r {
			s := nums[i] + nums[l] + nums[r]
			if s > target {
				r--
			} else if s < target {
				l++
			} else {
				l++
				r--
				return target
			}
			diff := abs(s - target)
			if diff < min {
				min = diff
				sum = s
			}
		}
	}
	return sum
}
func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}
