package applyoperationtoanarray

func applyOperations(nums []int) []int {
	result := make([]int, len(nums))
	k := 0
	i := 0
	for ; i < len(nums)-1; i++ {
		if nums[i] == nums[i+1] {
			nums[i] *= 2
			nums[i+1] = 0
		}

		if nums[i] > 0 {
			result[k] = nums[i]
			k++
		}
	}
	if nums[i] > 0 {
		result[k] = nums[i]
		k++
	}
	return result
}
