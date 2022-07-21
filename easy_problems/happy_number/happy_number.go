package happy_number

func isHappyNum(n int) bool {
	if n == 1 || n == 7 {
		return true
	}
	for n > 9 {
		sum := 0

		for n > 0 {
			x := n % 10
			sum += x * x
			n /= 10
		}
		n = sum
	}
	if n == 1 || n == 7 {
		return true
	}
	return false
}
