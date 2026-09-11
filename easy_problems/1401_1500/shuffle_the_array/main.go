package shufflethearray

func shuffle(nums []int, n int) []int {
	ans := make([]int, 2*n)
	i, j := 0, n
	for k := 0; i < n && j < 2*n; i, j, k = i+1, j+1, k+2 {
		ans[k], ans[k+1] = nums[i], nums[j]
	}
	return ans
}
