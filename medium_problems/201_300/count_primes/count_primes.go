package count_primes

import "math"

func CountPrimes(n int) int {
	result := make([]bool, n+1)
	for i := 0; i <= n; i++ {
		result[i] = true
	}
	m := int(math.Sqrt(float64(n)))
	for i := 2; i <= m; i++ {
		if result[i] {

			for j := i * i; j <= n; j = j + i {
				result[j] = false
			}
		}
	}
	count := 0

	for i := 2; i < n; i++ {
		if result[i] {
			count++
		}
	}

	return count
}
