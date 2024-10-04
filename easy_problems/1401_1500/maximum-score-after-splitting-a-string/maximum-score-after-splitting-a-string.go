package maximumscoreaftersplittingastring

func maxScore(s string) int {
	zero, one := make([]int, len(s)), make([]int, len(s))

	if s[0] == '0' {
		zero[0]++
	} else {
		one[0]++
	}
	for i := 1; i < len(s); i++ {
		one[i] = one[i-1]
		zero[i] = zero[i-1]

		if s[i] == '0' {
			zero[i]++
		} else {
			one[i]++
		}
	}
	if zero[len(s)-1] == len(s) || one[len(s)-1] == len(s) {
		return len(s) - 1
	}
	max := 0
	for i := 1; i < len(s); i++ {
		sum := zero[i-1] + one[len(s)-1] - one[i-1]
		if sum > max {
			max = sum
		}
	}
	return max
}
