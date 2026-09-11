package check_if_there_valid_partition_for_the_array

// ValidPartition reports whether nums can be split into contiguous
// pieces that are each two equal elements, three equal elements, or
// three consecutive increasing elements.
//
// A piece can only end at one of three offsets behind the current
// position, so whether a prefix is splittable depends on at most two
// shorter prefixes. dp[i] is that answer for the first i elements.
func ValidPartition(nums []int) bool {
	n := len(nums)
	dp := make([]bool, n+1)
	dp[0] = true

	for i := 2; i <= n; i++ {
		if dp[i-2] && nums[i-1] == nums[i-2] {
			dp[i] = true
			continue
		}
		if i < 3 || !dp[i-3] {
			continue
		}
		a, b, c := nums[i-3], nums[i-2], nums[i-1]
		if (a == b && b == c) || (b == a+1 && c == b+1) {
			dp[i] = true
		}
	}
	return dp[n]
}
