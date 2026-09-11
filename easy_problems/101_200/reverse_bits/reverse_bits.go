package reverse_bits

func ReverseBits(num uint32) uint32 {
	sum := uint32(0)
	for i := 31; i >= 0; i-- {
		x := uint32(1 << i)
		sum += (num % 2) * x
		num /= 2
	}
	return sum
}
