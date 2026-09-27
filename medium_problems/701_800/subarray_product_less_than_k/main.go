package subarrayproductlessthank

func numSubarrayProductLessThanK(nums []int, k int) int {
	var pro, ka int64 = 1, int64(k)
	count := 0
	left, right, n := 0, 0, len(nums)
	for right < n {
		pro *= int64(nums[right])
		// Values are >= 1, so dropping from the left never raises the product.
		// Division is exact: every value divided out was multiplied in.
		for left <= right && pro >= ka {
			pro /= int64(nums[left])
			left++
		}
		// Every piece of a valid window is valid, so all right-left+1
		// subarrays ending at right count. For k <= 1 this is always 0.
		count += right - left + 1
		right++
	}
	return count
}
