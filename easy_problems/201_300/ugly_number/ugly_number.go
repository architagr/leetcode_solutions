package ugly_number

func IsUgly(n int) bool {
	if n == 0 {
		return false
	}

	for n%5 == 0 {
		n = n / 5
	}

	for n%3 == 0 {
		n = n / 3
	}

	for n%2 == 0 {
		n = n / 2
	}

	return n == 1
}
