package find_the_middle_index_in_array

func FindMiddleIndex(nums []int) int {
	prefixSum := getPrefixSum(nums)
	n := len(nums)

	for i := 0; i < n; i++ {
		right, left := 0, 0
		if i > 0 {
			left = prefixSum[i-1]
		}
		if i < n-1 {
			right = prefixSum[n-1] - prefixSum[i]
		}
		if right == left {
			return i
		}
	}
	return -1

}
func getPrefixSum(nums []int) []int {
	prefixSum := make([]int, len(nums))
	prefixSum[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		prefixSum[i] = prefixSum[i-1] + nums[i]
	}
	return prefixSum
}
