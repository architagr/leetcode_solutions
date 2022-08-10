package longest_substring_without_repeating_characters

func LengthOfLongestSubstring(s string) int {
	mapS := make(map[byte]int)
	maxLen := 0
	for i, j := 0, 0; i < len(s); i++ {
		if v, ok := mapS[s[i]]; ok {
			j = maxvalue(v, j)
		}
		maxLen = maxvalue(maxLen, i-j+1)
		mapS[s[i]] = i + 1
	}

	return maxLen
}

func maxvalue(a, b int) int {
	if a > b {
		return a
	}
	return b
}
