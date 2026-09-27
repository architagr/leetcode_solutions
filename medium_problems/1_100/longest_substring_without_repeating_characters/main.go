package longestsubstringwithoutrepeatingcharacters

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	count := 0
	start := 0
	// Letter -> index, for exactly the letters in the current window. That
	// makes len(uniqueChar) the window's length.
	uniqueChar := make(map[byte]int, 128)
	for i := 0; i < len(s); i++ {
		// A repeat: everything from start through the earlier copy must leave,
		// since a substring can't skip a letter. Each letter is deleted at
		// most once overall, so this stays O(n) in total.
		if r, found := uniqueChar[s[i]]; found {
			for ; start <= r; start++ {
				delete(uniqueChar, s[start])
			}
		}
		uniqueChar[s[i]] = i
		if len(uniqueChar) > count {
			count = len(uniqueChar)
		}
	}
	return count
}
