package word_subsets

func UniversalWordSubset(words1, words2 []string) []string {
	words2Map := ConvertWord2ToMap(words2)
	arr := make([]string, 0, len(words1))
	for _, word := range words1 {
		temp := convertWordToMap(word)
		i := 0
		for ; i < 26; i++ {
			if temp[i] < words2Map[i] {
				break
			}
		}
		if i == 26 {
			arr = append(arr, word)
		}
	}
	return arr
}

func ConvertWord2ToMap(words2 []string) []int {
	n := make([]int, 26)
	m := make([][]int, len(words2))
	for i, word := range words2 {
		m[i] = convertWordToMap(word)
		for j := 0; j < 26; j++ {
			n[j] = maxvalue(n[j], m[i][j])
		}
	}

	return n
}
func convertWordToMap(word string) []int {
	m := make([]int, 26)
	for i := 0; i < len(word); i++ {
		m[int(word[i]-'a')]++
	}
	return m
}
func maxvalue(a, b int) int {
	if a > b {
		return a
	}
	return b
}
