package wordpattern

import "strings"

func wordPattern(pattern string, s string) bool {
	m1 := make(map[byte]string, 26)
	m2 := make(map[string]byte, 26)
	words := strings.Split(s, " ")
	if len(words) != len(pattern) {
		return false
	}
	for i := 0; i < len(pattern); i++ {
		ch := pattern[i]
		word := words[i]
		mWord, m1found := m1[ch]
		mCh, m2found := m2[word]
		if m1found != m2found {
			return false
		}
		if m1found && (mWord != word || ch != mCh) {
			return false
		}
		m1[ch] = word
		m2[word] = ch
	}
	return true
}
