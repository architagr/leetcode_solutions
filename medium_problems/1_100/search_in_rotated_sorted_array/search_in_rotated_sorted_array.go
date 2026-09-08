package search_in_rotated_sorted_array

func Search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	start, end := 0, len(nums)-1

	for start <= end {
		mid := start + (end-start)/2
		if nums[mid] == target {
			return mid
		}
		if nums[start] <= nums[mid] {
			if target >= nums[start] && target <= nums[mid] {
				end = mid - 1
			} else {
				start = mid + 1
			}
		} else {
			if target >= nums[mid] && target <= nums[end] {
				start = mid + 1

			} else {
				end = mid - 1
			}
		}
	}
	return -1

}
