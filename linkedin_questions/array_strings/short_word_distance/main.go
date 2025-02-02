package shortworddistance

func shortestDistance(wordsDict []string, word1 string, word2 string) int {
	word1Count, word2Count := 1, 1
	result := 0
	for l, r := 0, 0; r < len(wordsDict); r++ {
		isWord1 := wordsDict[r] == word1
		isWord2 := wordsDict[r] == word2
		if !isWord1 && !isWord2 {
			continue
		}
		if isWord1 {
			word1Count--
		}
		if isWord2 {
			word2Count--
		}
		for ; word1Count <= 0 && word2Count <= 0 && l <= r; l++ {
			if result == 0 || result > r-l {
				result = r - l
			}
			isWord1L := wordsDict[l] == word1
			isWord2L := wordsDict[l] == word2
			if isWord1L {
				word1Count++
			}
			if isWord2L {
				word2Count++
			}
		}
	}
	return result
}
