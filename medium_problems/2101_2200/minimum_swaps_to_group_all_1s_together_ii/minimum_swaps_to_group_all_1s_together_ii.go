package minimumswapstogroupall1stogetherii

func minSwaps(nums []int) int {
	n, ans := len(nums), len(nums)
	maxOne := 0
	for _, val := range nums {
		if val == 1 {
			maxOne++
		}
	}
	if maxOne == 0 {
		return 0
	}
	// The grouped ones fill a block of length maxOne. Each 0 inside that
	// block needs exactly one swap with a 1 outside it, so the answer is the
	// fewest zeros in any circular window of that size.
	countZero := 0
	for i := 0; i < maxOne-1; i++ {
		if nums[i] == 0 {
			countZero++
		}
	}

	k := maxOne - 1
	// Every start position once; (i+k)%n lets windows run past the end.
	for i := 0; i < n; i++ {
		if nums[(i+k)%n] == 0 {
			countZero++
		}
		if ans > countZero {
			ans = countZero
		}
		if nums[i] == 0 {
			countZero--
		}
	}
	return ans
}
