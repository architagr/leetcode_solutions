package findfirstandlastpositionofelementinsortedarray

func searchRange(nums []int, target int) []int {
	index := searchSubArray(nums, 0, len(nums)-1, target)
	result := []int{index, index}

	for index = index + 1; index < len(nums) && nums[index] == target; index++ {
		result[1] = index
	}

	return result
}

func searchSubArray(nums []int, start, end int, target int) int {
	if len(nums) == 0 {
		return -1
	}
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
	if nums[m] >= target {
		// go left
		return searchSubArray(nums, start, m, target)
	} else if nums[m] < target {
		// go right
		return searchSubArray(nums, m+1, end, target)
	}
	return -1
}
