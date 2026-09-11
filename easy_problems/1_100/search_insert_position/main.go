package searchinsertposition

func searchInsert(nums []int, target int) int {
	end := len(nums) - 1
	if end < 0 {
		return 0
	} else {
		return searchData(nums, target)
	}
}
func searchData(nums []int, target int) int {
	start, end := 0, len(nums)-1
	mid := end / 2
	if end == start {
		if nums[start] >= target {
			return start
		} else {
			return start + 1
		}
	} else if nums[mid] == target {
		return start + mid
	} else if nums[mid] > target {
		return start + searchData(nums[start:mid+1], target)
	} else if nums[mid] < target {
		return mid + 1 + searchData(nums[mid+1:end+1], target)
	}

	return end
}
