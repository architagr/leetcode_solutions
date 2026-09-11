package separatethedigitsinanarray

func separateDigits(nums []int) []int {
	result := make([]int, 0, len(nums))
	reverse(nums)
	for _, n := range nums {
		result = append(result, digits(n)...)
	}
	reverse(result)
	return result
}

func digits(n int) []int {
	result := make([]int, 0, 5)
	for n > 0 {
		result = append(result, n%10)
		n /= 10
	}
	return result
}

func reverse(nums []int) {
	for i := 0; i < len(nums)/2; i++ {
		nums[i], nums[len(nums)-1-i] = nums[len(nums)-1-i], nums[i]
	}
}
