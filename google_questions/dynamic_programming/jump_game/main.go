package jumpgame

func canJump(nums []int) bool {
	goal := len(nums) - 1

	for i := goal; i >= 0; i-- {
		if i+nums[i] >= goal {
			goal = i
		}
	}
	return goal == 0
}
