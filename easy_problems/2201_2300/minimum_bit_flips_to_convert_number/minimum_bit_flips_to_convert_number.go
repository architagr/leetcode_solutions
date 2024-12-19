package minimumbitflipstoconvertnumber

func minBitFlips(start int, goal int) int {
	xor := start ^ goal
	count := 0
	for xor > 0 {
		if xor%2 == 1 {
			count++
		}
		xor /= 2
	}
	return count
}
