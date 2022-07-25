package power_of_four

func IsPowerOfFour(n int) bool {
	if n <= 0 {
		return false
	}
	if n == 1 {
		return true
	}

	for n > 1 {
		if n%4 != 0 {
			return false
		}
		n /= 4
	}
	return true
}
