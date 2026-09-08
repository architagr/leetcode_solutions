package longestsubstringwithoutrepeatingcharacters

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	count := 0
	start := 0
	uniqueChar := make(map[byte]int, 128)
	for i := 0; i < len(s); i++ {
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
