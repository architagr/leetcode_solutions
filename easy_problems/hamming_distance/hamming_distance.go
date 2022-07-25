package hamming_distance

func HammingDistance(x int, y int) int {

	count := 0
	for i := 0; i < 32; i++ {
		if (x % 2) != (y % 2) {
			count++
		}
		x /= 2
		y /= 2
	}

	return count
}
