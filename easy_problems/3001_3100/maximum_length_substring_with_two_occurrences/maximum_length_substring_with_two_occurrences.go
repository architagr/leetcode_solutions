package maximumlengthsubstringwithtwooccurrences

func maximumLengthSubstring(s string) int {

	frequencyMap := make(map[byte]int, 26)
	m := 0
	left, right := 0, 0
	for ; right < len(s); right++ {
		frequencyMap[s[right]]++
		if frequencyMap[s[right]] == 3 {
			// The window only grows between violations, so s[left:right] is
			// the longest it got; record it before shrinking.
			m = max(right-left, m)
			// Shrink past the first copy of s[right]. Everything before that
			// copy has to go too, since a substring can't skip letters.
			for ; left <= right; left++ {
				frequencyMap[s[left]]--
				if s[right] == s[left] {
					left++
					break
				}
			}
		}
	}
	// The final window never hit a violation, so it hasn't been recorded yet.
	m = max(right-left, m)
	return m
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
