package xor_queries_of_a_subarray

func XorQueries(arr []int, queries [][]int) []int {
	prefixSum := getPrefixSum(arr)
	result := make([]int, len(queries))
	for i := 0; i < len(queries); i++ {
		result[i] = prefixSum[queries[i][1]]
		if queries[i][0] > 0 {
			result[i] ^= prefixSum[queries[i][0]-1]
		}
	}
	return result
}
func getPrefixSum(arr []int) []int {
	prefixSum := make([]int, len(arr))
	prefixSum[0] = arr[0]
	for i := 1; i < len(arr); i++ {
		prefixSum[i] = prefixSum[i-1] ^ arr[i]
	}
	return prefixSum
}
