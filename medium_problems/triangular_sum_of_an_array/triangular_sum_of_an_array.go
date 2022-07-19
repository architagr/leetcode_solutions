package triangular_sum_of_an_array

func TriangularSum(nums []int) int {
	for len(nums) > 1 {
		for i := 0; i < len(nums)-1; i++ {
			nums[i] = (nums[i] + nums[i+1]) % 10
		}
		nums = nums[:len(nums)-1]
	}
	return nums[0]
}

func triangularSum(nums []int) int {
	n := len(nums)

	if n == 1 {
		return nums[0]
	}

	for i := 0; i < n-1; i++ {
		nums[i] = (nums[i] + nums[i+1]) % 10
	}
	return triangularSum(nums[:n-1])
}
