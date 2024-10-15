package minimumoperationstomakebinaryarrayelementsequaltoonei

func minOperations(nums []int) int {
	count, k := 0, 3

	for i := 0; i <= len(nums)-k; i++ {
		if nums[i] == 0 {
			count++
			nums[i] = 1
			if nums[i+1] == 0 {
				nums[i+1] = 1
			} else {
				nums[i+1] = 0
			}
			if nums[i+2] == 0 {
				nums[i+2] = 1
			} else {
				nums[i+2] = 0
			}
		}
	}
	for i := 0; i <= len(nums)-1; i++ {
		if nums[i] == 0 {
			return -1
		}
	}
	return count
}
