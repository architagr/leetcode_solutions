package movezeroes

func moveZeroes(nums []int) {
	i, j := 0, 0
	for i < len(nums) && nums[i] != 0 {
		i++
		j++
	}
	for j < len(nums) {
		if nums[i] == 0 {
			if nums[j] != 0 {
				nums[i], nums[j] = nums[j], nums[i]
				i++
			} else {
				j++
			}
		}
	}
}
