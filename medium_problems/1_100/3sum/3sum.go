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

			for nums[l-1] == nums[l] && l < r {
				l++
			}

			for nums[r+1] == nums[r] && l < r {
				r--
			}
		}
	}
	return result
}
