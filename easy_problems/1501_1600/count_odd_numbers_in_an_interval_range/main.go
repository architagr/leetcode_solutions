package countoddnumbersinanintervalrange

func countOdds(low int, high int) int {
	if low&1 != 1 {
		low++
	}
	if high&1 != 1 {
		high--
	}
	return ((high - low) / 2) + 1
}
