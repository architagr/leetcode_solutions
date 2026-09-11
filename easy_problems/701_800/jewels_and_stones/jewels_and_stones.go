package jewels_and_stones

func CountJewels(jewels, stones string) int {
	m := make(map[byte]bool)
	count := 0
	for i := 0; i < len(jewels); i++ {
		m[jewels[i]] = true
	}
	for i := 0; i < len(stones); i++ {
		if _, ok := m[stones[i]]; ok {
			count++
		}
	}

	return count
}
