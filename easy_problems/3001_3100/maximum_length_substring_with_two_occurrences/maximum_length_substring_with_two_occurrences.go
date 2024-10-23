package maximumlengthsubstringwithtwooccurrences

func maximumLengthSubstring(s string) int {

	frequencyMap := make(map[byte]int, 26)
	m := 0
	left, right := 0, 0
	for ; right < len(s); right++ {
		frequencyMap[s[right]]++
		if frequencyMap[s[right]] == 3 {
			m = max(right-left, m)
			for ; left <= right; left++ {
				frequencyMap[s[left]]--
				if s[right] == s[left] {
					left++
					break
				}
			}
		}
	}
	m = max(right-left, m)
	return m
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
