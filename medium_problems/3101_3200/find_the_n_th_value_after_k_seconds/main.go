package findthenthvalueafterkseconds

func valueAfterKSeconds(n int, k int) int {
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = 1
	}
	mod := 1000000007
	for k > 0 {

		for i := 1; i < n; i++ {
			arr[i] = (arr[i-1] + arr[i]) % mod
		}
		k--
	}
	return arr[n-1]
}
