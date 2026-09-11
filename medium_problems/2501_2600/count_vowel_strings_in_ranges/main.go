package countvowelstringsinranges

func vowelStrings(words []string, queries [][]int) []int {
	count := 0
	nW := len(words)
	arr := make([]int, nW)

	for i, word := range words {
		n := len(word)
		a := 0
		if isVowel(word[0]) && isVowel(word[n-1]) {
			a = 1
		}
		count += a
		arr[i] = count
	}

	result := make([]int, len(queries))

	for i, q := range queries {
		result[i] = arr[q[1]]
		if q[0] > 0 {
			result[i] -= arr[q[0]-1]
		}
	}
	return result
}

func isVowel(c byte) bool {
	return c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u'
}
