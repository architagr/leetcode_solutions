package countnumberofnicesubarrays

func numberOfSubarrays(nums []int, k int) int {
	n := len(nums)
	ans, prefixSum := 0, 0
	// Odd -> 1, even -> 0 makes this "subarrays summing to exactly k".
	// m[c] is how many prefixes so far had c odd numbers; seeding m[0] is
	// the empty prefix, so subarrays starting at index 0 are counted.
	m := make(map[int]int)
	m[0]++
	for i := 0; i < n; i++ {
		if nums[i]%2 == 1 {
			prefixSum++
		}
		// Every earlier prefix with k fewer odds starts a nice subarray here.
		ans += m[prefixSum-k]
		m[prefixSum]++
	}
	return ans
}
