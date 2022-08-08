package longest_ideal_subsequence

func LongestIdealString(s string, k int) int {
	length := 0

	result := ""
	result += string(s[len(s)-1])

	for i := len(s) - 2; i >= 0; i-- {
		diff := int(result[0]) - int(s[i])
		if diff < 0 {
			diff *= -1
		}
		if diff <= k {
			result = string(s[i]) + result
		}
	}
	if len(result) == 1 {
		length = 0
	} else {
		length = len(result)
	}

	return length
}
