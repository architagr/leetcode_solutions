package validword

func isValid(word string) bool {
	if len(word) < 3 {
		return false
	}

	hasVowel, hasConsonant := false, false

	for i := 0; i < len(word); i++ {
		c := byte(word[i])
		if c >= 48 && c <= 57 {
			continue
		} else if c >= 65 && c <= 90 {
			if c == 65 || c == 69 || c == 73 || c == 79 || c == 85 {
				hasVowel = true
			} else {
				hasConsonant = true
			}
			continue
		} else if c >= 97 && c <= 122 {
			if c == 97 || c == 101 || c == 105 || c == 111 || c == 117 {
				hasVowel = true
			} else {
				hasConsonant = true
			}
			continue
		} else {
			return false
		}
	}
	return hasVowel && hasConsonant
}
