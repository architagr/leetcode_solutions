package number_of_matching_subsequences

func NumMatchingSubseq(s string, words []string) int {
	wordmap := make(map[string]int)
	for i := 0; i < len(words); i++ {
		wordmap[words[i]]++
	}
	ans := 0
	for word, count := range wordmap {
		i, j := 0, 0
		for i < len(s) && j < len(word) {
			if s[i] == word[j] {
				i++
				j++
			} else {
				i++
			}
		}
		if j == len(word) {
			ans += count
		}
	}
	return ans
}
