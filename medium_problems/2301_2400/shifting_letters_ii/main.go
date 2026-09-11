package shiftinglettersii

func shiftingLetters(s string, shifts [][]int) string {
	n := len(s)
	arr := make([]int, n)
	for _, shift := range shifts {
		factor := 1
		if shift[2] == 0 {
			factor = -1
		}
		arr[shift[0]] += factor
		if shift[1] < n-1 {
			arr[shift[1]+1] += (-1 * factor)
		}
	}
	result := make([]byte, n)
	sum := 0
	for i := 0; i < n; i++ {
		sum += arr[i]
		result[i] = 'a' + byte((((int(s[i]-'a')+sum)%26)+26)%26)
	}

	return string(result)
}
