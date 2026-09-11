package binary_number_with_alternating_bits

func HasAlternatingBits(n int) bool {
	x := n % 2
	n /= 2
	for n > 0 {
		if x == n%2 {
			return false
		}
		x = n % 2
		n /= 2
	}
	return true
}
