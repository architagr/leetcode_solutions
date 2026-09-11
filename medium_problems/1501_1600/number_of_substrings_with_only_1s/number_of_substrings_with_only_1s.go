package number_of_substrings_with_only_1s

func NumSub(s string) int {
	s += "0"
	n := len(s)

	count, sum := 0, 0
	for i := 0; i < n; i++ {
		if s[i] == '1' {
			count++
		} else if count > 0 {
			count = count * (count + 1) / 2
			sum += count % 1000000007
			count = 0
		}
	}

	return sum % 1000000007
}
