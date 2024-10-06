package sentencesimilarityiii

import "strings"

func areSentencesSimilar(sentence1 string, sentence2 string) bool {
	// Split sentences into words
	s1 := strings.Split(sentence1, " ")
	s2 := strings.Split(sentence2, " ")

	// Ensure s2 is the longer sentence
	if len(s1) > len(s2) {
		s1, s2 = s2, s1
	}

	left, right := 0, 0
	n1, n2 := len(s1), len(s2)

	// Compare from the start
	for left < n1 && s1[left] == s2[left] {
		left++
	}

	// Compare from the end
	for right < n1 && s1[n1-right-1] == s2[n2-right-1] {
		right++
	}

	// Check if the remaining unmatched part is in the middle
	return left+right >= n1
}
