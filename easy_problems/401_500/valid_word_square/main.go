package validwordsquare

func validWordSquare(words []string) bool {
	b := make([][]byte, len(words))
	max := len(words)

	for i, word := range words {
		b[i] = []byte(word)
		max = maxValue(max, len(b[i]))
	}
	for i := 0; i < len(b); i++ {
		if len(b[i]) < max {
			x := make([]byte, max-len(b[i]))
			b[i] = append(b[i], x...)
		}
	}
	for i := len(b); i < max; i++ {
		b = append(b, make([]byte, max))
	}

	for i := 0; i < max; i++ {
		for j := i; j < max; j++ {
			if b[i][j] != b[j][i] {
				return false
			}
		}
	}
	return true
}

func maxValue(a, b int) int {
	if a > b {
		return a
	}
	return b
}
