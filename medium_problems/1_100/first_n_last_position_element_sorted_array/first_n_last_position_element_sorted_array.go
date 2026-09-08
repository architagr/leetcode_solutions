package first_n_last_position_element_sorted_array

func SearchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}
	start := binarySearch(nums, target, true)
	end := binarySearch(nums, target, false)
	return []int{start, end}
}

func binarySearch(nums []int, target int, low bool) int {
	start, end := 0, len(nums)-1
	ans := -1
	for start <= end {
		mid := start + (end-start)/2
		if nums[mid] == target {
			ans = mid
			if low {
				end = mid - 1
			} else {
				start = mid + 1
			}
		} else if nums[mid] > target {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}
	return ans
}
