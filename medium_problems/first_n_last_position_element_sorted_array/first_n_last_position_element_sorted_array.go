package first_n_last_position_element_sorted_array

func SearchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}
	left := binarySearch(nums, target, true)
	right := binarySearch(nums, target, false)
	return []int{left, right}
}

func binarySearch(nums []int, target int, low bool) int {
	left, right := 0, len(nums)-1
	ans := -1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			ans = mid
			if low {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return ans
}
