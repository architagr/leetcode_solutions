package find_the_duplicate_number

func findDuplicate(nums []int) int {
	slow, fast := 0, 0
	for {
		slow = nums[slow]
		fast = nums[nums[fast]]
		if slow == fast {
			break
		}
	}
	slow2 := 0
	for {
		slow = nums[slow]
		slow2 = nums[slow2]
		if slow == slow2 {
			break
		}
	}
	return slow
}

//https://www.youtube.com/watch?v=wjYnzkAhcNk
