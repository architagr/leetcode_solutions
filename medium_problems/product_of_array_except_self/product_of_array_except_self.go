package product_of_array_except_self

func ProductExceptSelf(nums []int) []int {
	left := getLeftMultiplicaion(nums)
	right := getRightMultiplicaion(nums)

	nums[0] = right[1]
	nums[len(nums)-1] = left[len(nums)-2]
	for i := 1; i < len(nums)-1; i++ {
		nums[i] = left[i-1] * right[i+1]
	}
	return nums
}

func getLeftMultiplicaion(nums []int) []int {
	arr := make([]int, len(nums))
	mul := 1
	for i := 0; i < len(nums)-1; i++ {
		mul *= nums[i]
		arr[i] = mul
	}
	return arr
}

func getRightMultiplicaion(nums []int) []int {
	arr := make([]int, len(nums))
	mul := 1
	for i := len(nums) - 1; i >= 0; i-- {
		mul *= nums[i]
		arr[i] = mul
	}
	return arr
}
