package maxconsecutiveonesiii

func longestOnes(nums []int, k int) int {
	l, r := 0, 0
	count := 0
	ans := 0
	// Flipping at most k zeros is the same as finding the longest window with
	// at most k zeros; count is the number of zeros in [l, r].
	for ; r < len(nums); r++ {
		// A one can never cause a violation, so skip the rest.
		if nums[r] == 1 {
			continue
		}
		count++
		for ; l <= r && l < len(nums) && count > k; l++ {
			// r-l is [l, r-1], the window before this zero arrived; the first
			// pass through here is the only one that can raise ans.
			if ans < r-l {
				ans = r - l
			}
			if nums[l] == 0 {
				count--
			}
		}
	}
	// The final window may never have hit a violation.
	if count <= k && ans < r-l {
		ans = r - l
	}
	return ans
}
