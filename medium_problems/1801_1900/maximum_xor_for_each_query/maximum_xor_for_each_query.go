package maximum_xor_for_each_query

func GetMaximumXor(nums []int, maximumBit int) []int {
	n := len(nums)
	prefixSum := make([]int, n)
	prefixSum[0] = nums[0]
	for i := 1; i < n; i++ {
		prefixSum[i] = prefixSum[i-1] ^ nums[i]
	}
	x := (1 << maximumBit) - 1
	result := make([]int, n)
	for i := 0; i < n; i++ {
		result[i] = prefixSum[n-1-i] ^ x
	}
	return result
}
