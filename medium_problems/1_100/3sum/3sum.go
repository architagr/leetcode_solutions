package sum3

import "sort"

func ThreeSum(nums []int) [][]int {
	result := make([][]int, 0)
	n := len(nums)
	sort.Ints(nums)

	for i := 0; i < n; i++ {
		if i > 0 && nums[i-1] == nums[i] {
			continue
		}

		l, r := i+1, n-1
		for l < r {
			sum := nums[i] + nums[l] + nums[r]
			if sum > 0 {
				r--
			} else if sum < 0 {
				l++
			} else {
				result = append(result, []int{nums[i], nums[l], nums[r]})
				l++
				r--
			}

			// Both guards have to come before the index, or l runs past
			// the end of nums and r+1 reads one past it on the first pass.
			for l < r && nums[l-1] == nums[l] {
				l++
			}

			for l < r && r+1 < n && nums[r+1] == nums[r] {
				r--
			}
		}
	}
	return result
}
