package partition_array_according_given_pivot

func pivotArray(nums []int, pivot int) []int {
	ans := make([]int, len(nums))
	l, r := 0, len(nums)-1

	for i := 0; i < len(nums); i++ {
		ans[i] = pivot

		if nums[i] < pivot {
			ans[l] = nums[i]
			l++
		}
	}

	for i := 0; i < len(nums); i++ {
		if nums[len(nums)-1-i] > pivot {
			ans[r] = nums[len(nums)-1-i]
			r--
		}
	}

	return ans
}
