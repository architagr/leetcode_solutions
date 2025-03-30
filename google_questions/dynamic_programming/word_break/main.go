package wordbreak

func wordBreak(s string, wordDict []string) bool {
	n := len(s)
	wordCheck := make([]bool, n+1)
	wordCheck[0] = true

	for i := 0; i < n; i++ {
		for _, word := range wordDict {
			l := len(word)
			if i-l+1 >= 0 && s[i-l+1:i+1] == word && wordCheck[i-l+1] {
				wordCheck[i+1] = true
			}
		}
	}
	return wordCheck[n]
}
