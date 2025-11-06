package missingranges

func findMissingRanges(nums []int, lower int, upper int) [][]int {
	result := make([][]int, 0)
	nums = append([]int{lower - 1}, nums...) // this is the key logic, add lower - 1 at beginning
	nums = append(nums, upper+1)             //append upper + 1 at the end

	for i := 0; i < len(nums)-1; i++ {
		l := nums[i] + 1
		r := nums[i+1] - 1
		if l <= r {
			result = append(result, []int{l, r})
		}
	}
	return result
}
