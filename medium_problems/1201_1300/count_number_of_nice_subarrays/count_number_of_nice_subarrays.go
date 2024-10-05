package countnumberofnicesubarrays

func numberOfSubarrays(nums []int, k int) int {
	n := len(nums)
	ans, prefixSum := 0, 0
	m := make(map[int]int)
	m[0]++
	for i := 0; i < n; i++ {
		if nums[i]%2 == 1 {
			prefixSum++
		}
		ans += m[prefixSum-k]
		m[prefixSum]++
	}
	return ans
}
