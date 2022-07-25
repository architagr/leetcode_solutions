package target_Indices_after_sorting_array

import "sort"

func TargetIndices(nums []int, target int) []int {
	sort.Ints(nums)
	if len(nums) == 0 {
		return []int{-1, -1}
	}
	start := binarySearch(nums, target, true)
	end := binarySearch(nums, target, false)
	arr := make([]int, 0, len(nums))
	if start > -1 {
		for i := start; i <= end; i++ {
			arr = append(arr, i)
		}
	}
	return arr
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
