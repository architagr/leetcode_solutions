package continuous_subarray_sum

func CheckSubarraySum(nums []int, k int) bool {
	mapR := make(map[int]int)
	sum := 0
	if len(nums) < 2 {
		return false
	}
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		sum = sum % k
		if sum == 0 && i > 0 {
			return true
		}
		//here we are checking that sum contains sum and length is greter than one
		if val, ok := mapR[sum]; ok && i-val > 1 {
			return true
		}

		if _, ok := mapR[sum]; !ok {
			mapR[sum] = i
		}
	}
	return false
}
