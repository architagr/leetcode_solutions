package xoroperationinanarray

func xorOperation(n int, start int) int {
	ans := 0

	for i := 0; i < n; i++ {
		ans = ans ^ (start + (2 * i))
	}
	return ans
}
