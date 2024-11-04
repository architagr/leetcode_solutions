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
	countZero := 0
	for i := 0; i < maxOne-1; i++ {
		if nums[i] == 0 {
			countZero++
		}
	}

	k := maxOne - 1
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
