package count_vowel_substrings

func CountVowelSubstrings(word string) int {
	vowelMap := map[byte]int{
		'a': 1,
		'e': 2,
		'i': 4,
		'o': 8,
		'u': 16,
	}
	count := 0
	for i := 0; i < len(word); i++ {
		if !validChar(word[i]) {
			continue
		}
		for x, j := 0, i; j < len(word); j++ {
			if validChar(word[j]) {
				x = x | vowelMap[word[j]]
				if x == 31 {
					count++
				}
			} else {
				break
			}
		}
	}
	return count
}
func validChar(ch byte) bool {
	return ch == 'a' || ch == 'e' || ch == 'i' || ch == 'o' || ch == 'u'
}
