package find_and_replace_pattern

func FindAndReplacePattern(words []string, pattern string) []string {
	wordsArr := make([]string, 0, len(words))

	for _, word := range words {
		if len(word) == len(pattern) {
			if CheckEachWord(word, pattern) {
				wordsArr = append(wordsArr, word)
			}
		}
	}
	return wordsArr
}

func CheckEachWord(word, pattern string) bool {
	m := make(map[byte]byte)
	s := make(map[byte]bool)
	for i := 0; i < len(pattern); i++ {
		if val, ok := m[pattern[i]]; ok {
			if val != word[i] {
				return false
			}
		} else {
			if _, ok := s[word[i]]; ok {
				return false
			}
			m[pattern[i]] = word[i]
			s[word[i]] = true
		}

	}
	return true
}
