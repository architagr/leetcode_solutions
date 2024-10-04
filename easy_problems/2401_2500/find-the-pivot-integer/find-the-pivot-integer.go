package findthepivotinteger

func pivotInteger(n int) int {
	last := (n * (n + 1))
	for i := 1; i <= n; i++ {
		if 2*i*i == last {
			return i
		}
	}
	return -1
}
