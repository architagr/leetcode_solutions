package minimumdifferencebetweenhighestandlowestofkscores

import "sort"

func minimumDifference(nums []int, k int) int {
	// Position in the input doesn't matter, so sort: after sorting, the best
	// group of k is always a run of neighbours. Note this sorts the caller's
	// slice in place.
	sort.Ints(nums)
	n := len(nums)
	ans := nums[n-1] - nums[0]
	k--
	for i := k; i < n; i++ {
		// In a sorted window the min is the first element and the max the last,
		// so there's no running state to keep: just read the two ends.
		l := nums[i] - nums[i-k]
		if l < ans {
			ans = l
		}
	}
	return ans
}
