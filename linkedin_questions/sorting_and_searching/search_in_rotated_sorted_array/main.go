package searchinrotatedsortedarray

func search(nums []int, target int) int {
	pivot := findPivotIndex(nums)
	if pivot == 0 {
		return searchSubArray(nums, 0, len(nums)-1, target)
	}
	if target < nums[0] {
		// go to right side
		return searchSubArray(nums, pivot, len(nums)-1, target)
	}
	return searchSubArray(nums, 0, pivot-1, target)
}

func findPivotIndex(nums []int) int {
	for i := 1; i < len(nums); i++ {
		if nums[i] < nums[i-1] {
			return i
		}
	}
	return 0
}
func searchSubArray(nums []int, start, end int, target int) int {
	if start == end {
		if nums[start] == target {
			return start
		}
		return -1
	}
	if end-start == 1 {
		if nums[start] == target {
			return start
		}
		if nums[end] == target {
			return end
		}
		return -1
	}
	m := (start + end) / 2
	if nums[m] > target {
		// go left
		return searchSubArray(nums, start, m-1, target)
	} else if nums[m] < target {
		// go right
		return searchSubArray(nums, m+1, end, target)
	} else if nums[m] == target {
		return m
	}
	return -1
}
