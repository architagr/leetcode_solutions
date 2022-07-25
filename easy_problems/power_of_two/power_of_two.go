package power_of_two

func IsPowerOfTwo(n int) bool {
	for i := 0; i <= 31; i++ {
		if (1 << i) == n {
			return true
		}
	}
	return false
}
