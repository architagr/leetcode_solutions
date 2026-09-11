package find_all_numbers_disappeared_in_an_array

func FindDisappearedNumbers(nums []int) []int {
	n := len(nums)
	result := make([]int, 0)
	for i := 0; i < n; i++ {
		x := nums[i]
		if nums[i] > n {
			x -= n
		}
		if nums[x-1] <= n {
			nums[x-1] += n
		}
	}
	for i := 0; i < n; i++ {
		if nums[i] <= n {
			result = append(result, i+1)
		}
	}

	return result
}
