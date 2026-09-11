package first_missing_positive

func FirstMissingPositiveNumber(nums []int) int {
	i := 0
	for i < len(nums) {
		if nums[i] >= 1 && nums[i] != i+1 && nums[i] <= len(nums) && nums[i] != nums[nums[i]-1] {
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
		} else {
			i++
		}
	}

	for i := 0; i < len(nums); i++ {
		if nums[i] != i+1 {
			return i + 1
		}
	}
	return len(nums) + 1
}
