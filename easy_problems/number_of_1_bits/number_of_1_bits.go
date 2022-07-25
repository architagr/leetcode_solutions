package number_of_1_bits

func HammingWeight(num uint32) int {
	sum := 0
	for num > 0 {
		if num%2 == 1 {
			sum++
		}
		num /= 2
	}
	return sum
}
