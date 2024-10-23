package minimumdifferencebetweenhighestandlowestofkscores

import "sort"

func minimumDifference(nums []int, k int) int {
	sort.Ints(nums)
	n := len(nums)
	ans := nums[n-1] - nums[0]
	k--
	for i := k; i < n; i++ {
		l := nums[i] - nums[i-k]
		if l < ans {
			ans = l
		}
	}
	return ans
}
