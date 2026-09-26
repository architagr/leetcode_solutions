package maximumnumberofvowelsinasubstringofgivenlength

func maxVowels(s string, k int) int {
	max := 0
	count := 0
	// k becomes an offset: s[i-k] is the oldest letter in the window.
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
		// A window can hold at most k vowels, so reaching k here could end
		// the loop early; it isn't needed for correctness.
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
