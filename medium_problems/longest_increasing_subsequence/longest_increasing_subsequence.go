package longest_increasing_subsequence

func LengthOfLIS(nums []int) int {
	n := len(nums)
	lis := make([]int, n)
	max := 1
	for i := n - 1; i >= 0; i-- {
		lis[i] = 1
		for j := i + 1; j < n; j++ {
			if nums[i] < nums[j] {
				lis[i] = MaxValue(lis[i], 1+lis[j])
			}
		}
		max = MaxValue(max, lis[i])
	}
	return max
}
func MaxValue(a, b int) int {
	if a > b {
		return a
	}
	return b
}
