package antontheboundary

func returnToBoundaryCount(nums []int) int {
	count, sum := 0, 0

	for _, step := range nums {
		sum += step
		if sum == 0 {
			count++
		}
	}

	return count
}
