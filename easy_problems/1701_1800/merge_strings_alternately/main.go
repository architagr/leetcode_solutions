package mergestringsalternatively

import "strings"

func mergeAlternately(word1 string, word2 string) string {
	res := strings.Builder{}

	wPtr1 := 0
	wPtr2 := 0

	for wPtr1 < len(word1) || wPtr2 < len(word2) {
		if wPtr1 < len(word1) {
			res.WriteByte(word1[wPtr1])
			wPtr1++
		}

		if wPtr2 < len(word2) {
			res.WriteByte(word2[wPtr2])
			wPtr2++
		}
	}

	return res.String()
}
