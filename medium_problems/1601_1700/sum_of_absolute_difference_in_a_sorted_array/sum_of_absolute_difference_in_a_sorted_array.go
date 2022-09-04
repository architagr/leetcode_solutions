package sum_of_absolute_difference_in_a_sorted_array

func GetSumAbsoluteDifferences(nums []int) []int {
	n := len(nums)
	prefixSum := make([]int, n)
	arr := make([]int, n)
	prefixSum[0] = nums[0]
	for i := 1; i < n; i++ {
		prefixSum[i] = prefixSum[i-1] + nums[i]
	}

	for i := 0; i < n; i++ {
		x := (nums[i] * (i + 1)) - prefixSum[i]
		y := prefixSum[n-1] - prefixSum[i] - (nums[i] * (n - i - 1))
		arr[i] = x + y
		if arr[i] < 0 {
			arr[i] *= -1
		}
	}
	return arr
}
