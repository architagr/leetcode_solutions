package shiftingletters

func shiftingLetters(s string, shifts []int) string {
	byteArr := make([]byte, len(s))
	pre := prefixSum(shifts)

	for i := 0; i < len(s); i++ {
		byteArr[i] = byte('a' + (int(s[i]-'a')+pre[i])%26)
	}
	return string(byteArr)
}

func prefixSum(n []int) []int {
	l := len(n)
	result := make([]int, l)
	result[l-1] = n[l-1]
	for i := l - 2; i >= 0; i-- {
		result[i] = result[i+1] + n[i]
	}
	return result
}
