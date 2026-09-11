package removeelement

func removeElement(nums []int, val int) int {
	for i := 0; i < len(nums)-1; {
		if nums[i] == val {
			nums = append(nums[:i], nums[i+1:]...)
		} else {
			i++
		}
	}
	length := len(nums)
	if length > 0 && nums[len(nums)-1] == val {
		nums = nums[:length-1]
	}
	return len(nums)
}
