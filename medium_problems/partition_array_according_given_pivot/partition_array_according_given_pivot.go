package partition_array_according_given_pivot

func PivotArray(nums []int, pivot int) []int {
	n := len(nums)
	left := make([]int, 0, n)
	arr := make([]int, 0, n)
	right := make([]int, 0, n)

	for i := 0; i < n; i++ {
		if nums[i] < pivot {
			left = append(left, nums[i])
		} else if nums[i] == pivot {
			arr = append(arr, nums[i])
		} else {
			right = append(right, nums[i])
		}
	}
	return append(left, append(arr, right...)...)
}
