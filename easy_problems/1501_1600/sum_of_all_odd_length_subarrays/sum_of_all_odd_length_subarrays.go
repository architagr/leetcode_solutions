package sumofalloddlengthsubarrays

func SumOddLengthSubarrays(arr []int) int {
	prefixSum := findPerfixSum(arr)
	n := len(arr)
	ans := prefixSum[n-1]
	for i := 3; i <= n; i = i + 2 {
		for j := 0; j <= n-i; j++ {
			ans += prefixSum[j+i-1]
			if j > 0 {
				ans -= prefixSum[j-1]
			}
		}
	}
	return ans
}
func findPerfixSum(arr []int) []int {
	prefixSum := make([]int, len(arr))
	prefixSum[0] = arr[0]
	for i := 1; i < len(arr); i++ {
		prefixSum[i] = prefixSum[i-1] + arr[i]
	}
	return prefixSum
}
