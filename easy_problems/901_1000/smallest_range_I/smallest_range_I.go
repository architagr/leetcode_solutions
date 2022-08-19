package smallest_range_I

func SmallestRangeI(nums []int, k int) int {
	max := 0
	min := 10000
	for _, num := range nums {
		if num > max {
			max = num
		}
		if num < min {
			min = num
		}
	}
	score := (max - k) - (min + k)
	if score < 0 {
		return 0
	}
	return score
}
