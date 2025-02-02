package maxconsecutiveonesiii

func longestOnes(nums []int, k int) int {
	l, r := 0, 0
	count := 0
	ans := 0
	for ; r < len(nums); r++ {
		if nums[r] == 1 {
			continue
		}
		count++
		for ; l <= r && l < len(nums) && count > k; l++ {
			if ans < r-l {
				ans = r - l
			}
			if nums[l] == 0 {
				count--
			}
		}
	}
	if count <= k && ans < r-l {
		ans = r - l
	}
	return ans
}
