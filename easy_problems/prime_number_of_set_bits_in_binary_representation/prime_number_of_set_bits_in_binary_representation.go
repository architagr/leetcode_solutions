package prime_number_of_set_bits_in_binary_representation

import "math/bits"

func CountPrimeSetBits(left, right int) int {

	count := 0
	for i := left; i <= right; i++ {
		bitCount := bits.OnesCount(uint(i))
		if bitCount == 2 || bitCount == 3 || bitCount == 5 || bitCount == 7 || bitCount == 11 ||
			bitCount == 13 || bitCount == 17 || bitCount == 19 || bitCount == 23 || bitCount == 29 ||
			bitCount == 31 {
			count++
		}
	}
	return count
}
