package maximum_xor_for_each_query

func GetMaximumXor(nums []int, maximumBit int) []int {
	n := len(nums)
	y := 0
	for i := 1; i < n; i++ {
		y ^= nums[i]
	}
	x := (1 << maximumBit) - 1
	result := make([]int, n)
	for i := 0; i < n; i++ {
		result[i] = y ^ x
		y ^= nums[n-1-i]
	}
	return result
}
