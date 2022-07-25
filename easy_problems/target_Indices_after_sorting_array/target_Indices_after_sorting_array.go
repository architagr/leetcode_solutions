package target_Indices_after_sorting_array

import "sort"

func TargetIndices(nums []int, target int) []int {
	sort.Ints(nums)
	if len(nums) == 0 {
		return []int{-1, -1}
	}
	left := binarySearch(nums, target, true)
	right := binarySearch(nums, target, false)
	arr := make([]int, 0, len(nums))
	if left > -1 {
		for i := left; i <= right; i++ {
			arr = append(arr, i)
		}
	}
	return arr
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
