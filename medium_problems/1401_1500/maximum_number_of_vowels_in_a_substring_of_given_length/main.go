package maximumnumberofvowelsinasubstringofgivenlength

func maxVowels(s string, k int) int {
	max := 0
	count := 0
	k--

	for i := 0; i < k; i++ {
		if isVowel(s[i]) {
			count++
		}
	}
	for i := k; i < len(s); i++ {
		if isVowel(s[i]) {
			count++
		}
		if count > max {
			max = count
		}
		if isVowel(s[i-k]) {
			count--
		}
	}

	return max
}
func isVowel(b byte) bool {
	return b == 'a' || b == 'e' || b == 'i' || b == 'o' || b == 'u'
}
