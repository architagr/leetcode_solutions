package counting_bits

func CountBits(n int) []int {
	result := make([]int, n+1)

	for i := 1; i <= n; i++ {
		if i%2 == 0 {
			result[i] = result[i/2]
		} else {
			result[i] = 1 + result[i/2]
		}
	}
	return result
}
