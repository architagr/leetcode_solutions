package minimum_XOR_sum_of_two_arrays

import "math/bits"

// MinimumXORSum returns the smallest possible sum of nums1[i] ^ nums2[i]
// over every rearrangement of nums2.
//
// This is an assignment problem, not something a greedy pass can settle:
// the cheapest partner for one element can be the only good partner for
// another. n is at most 14, so the state is which elements of nums2 have
// been used. The number used is also how far into nums1 we are, so
// dp[mask] is the best cost of matching the first popcount(mask)
// elements of nums1 against exactly the elements of nums2 in mask.
func MinimumXORSum(nums1 []int, nums2 []int) int {
	n := len(nums2)
	const inf = int(^uint(0) >> 1)

	dp := make([]int, 1<<n)
	for i := range dp {
		dp[i] = inf
	}
	dp[0] = 0

	for mask := 0; mask < 1<<n; mask++ {
		if dp[mask] == inf {
			continue
		}
		i := bits.OnesCount(uint(mask))
		if i >= len(nums1) {
			continue
		}
		for j := 0; j < n; j++ {
			if mask&(1<<j) != 0 {
				continue
			}
			next := mask | 1<<j
			if cost := dp[mask] + (nums1[i] ^ nums2[j]); cost < dp[next] {
				dp[next] = cost
			}
		}
	}
	return dp[1<<n-1]
}
