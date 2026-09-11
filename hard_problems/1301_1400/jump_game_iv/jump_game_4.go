package jump_game_4

func JumpGame4(nums []int, k int) int {
	queue := []int{0}

	for i := 1; i < len(nums); i++ {
		if queue[0] < i-k {
			queue = queue[1:]
		}

		nums[i] = nums[i] + nums[queue[0]]
		for len(queue) > 0 && nums[queue[len(queue)-1]] <= nums[i] {
			queue = queue[:len(queue)-1]
		}

		queue = append(queue, i)
	}

	return nums[len(nums)-1]
}
