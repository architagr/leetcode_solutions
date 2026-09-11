package missing_number

func MissingNumber(nums []int) int {

	for i := 0; i < len(nums); {
		if nums[i] < len(nums) && nums[i] != i {
			nums[i], nums[nums[i]] = nums[nums[i]], nums[i]
		} else {
			i++
		}
	}
	for i := 0; i < len(nums); i++ {
		if i != nums[i] {
			return i
		}
	}
	return len(nums)
}
