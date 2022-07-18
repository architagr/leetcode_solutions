package find_all_duplicates_in_an_array

func FindDuplicates(nums []int) []int {
	n := len(nums)
	result := make([]int, 0)
	for i := 0; i < n; i++ {
		x := nums[i]
		if nums[i] > n {
			x -= n
		}
		if nums[x-1] > n {
			result = append(result, x)
		} else {
			nums[x-1] += n
		}
	}
	return result
}
