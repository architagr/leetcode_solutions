package make_array_zero_by_subtracting_equal_amounts

func MinimumOperations(nums []int) int {
	m := make(map[int]bool)

	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			m[nums[i]] = true
		}
	}
	return len(m)
}
