package set_mismatch

func FindErrorNums(nums []int) []int {
	count := make([]int, len(nums)+1)
	for i := 0; i < len(nums); i++ {
		count[nums[i]]++
	}

	res := []int{0, 0}
	for i := 0; i < len(count); i++ {
		if count[i] == 0 {
			res[1] = i
		} else if count[i] > 1 {
			res[0] = i
		}
	}
	return res
}
