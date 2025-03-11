package countofsubstringscontainingeveryvowelandkconsonantsii

func countOfSubstrings(word string, k int) int64 {
	numValidSubstrings := int64(0)
	start, end := 0, 0

	// keep track of counts of vowels and consonants
	vowelCount := make(map[byte]int, 5)
	consonantCount := 0

	// compute index of next consonant for all indices
	nextConsonant := make([]int, len(word))
	nextConsonantIndex := len(word)
	for i := len(word) - 1; i >= 0; i-- {
		nextConsonant[i] = nextConsonantIndex
		if !isVowel(word[i]) {
			nextConsonantIndex = i
		}
	}

	// start sliding window
	for end < len(word) {
		// insert new letter
		newLetter := word[end]

		// update counts
		if isVowel(newLetter) {
			vowelCount[newLetter]++
		} else {
			consonantCount++
		}

		// shrink window if too many consonants in our window
		for consonantCount > k {
			startLetter := word[start]
			if isVowel(startLetter) {
				vowelCount[startLetter]--
				if vowelCount[startLetter] == 0 {
					delete(vowelCount, startLetter)
				}
			} else {
				consonantCount--
			}
			start++
		}

		// while we have a valid window, try to shrink it
		for start < len(word) && len(vowelCount) == 5 && consonantCount == k {
			// count the current valid substring, as well as valid substrings produced by appending more vowels
			numValidSubstrings += int64(nextConsonant[end] - end)
			startLetter := word[start]
			if isVowel(startLetter) {
				vowelCount[startLetter]--
				if vowelCount[startLetter] == 0 {
					delete(vowelCount, startLetter)
				}
			} else {
				consonantCount--
			}

			start++
		}
		end++
	}

	return numValidSubstrings
}

func isVowel(c byte) bool {
	return c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u'
}
