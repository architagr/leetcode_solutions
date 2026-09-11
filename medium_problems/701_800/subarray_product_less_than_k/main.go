package subarrayproductlessthank

func numSubarrayProductLessThanK(nums []int, k int) int {
	var pro, ka int64 = 1, int64(k)
	count := 0
	left, right, n := 0, 0, len(nums)
	for right < n {
		pro *= int64(nums[right])
		for left <= right && pro >= ka {
			pro /= int64(nums[left])
			left++
		}
		count += right - left + 1
		right++
	}
	return count
}
