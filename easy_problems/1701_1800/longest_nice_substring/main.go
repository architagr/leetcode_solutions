package longestnicesubstring

func longestNiceSubstring(s string) string {
	var longest string

	for left := 0; left < len(s); left++ {
		var uppercase, lowercase [26]bool // Tracks the presence of uppercase and lowercase characters

		for right := left; right < len(s); right++ {
			if s[right] >= 'a' && s[right] <= 'z' {
				lowercase[s[right]-'a'] = true
			} else {
				uppercase[s[right]-'A'] = true
			}

			// Check if the substring is nice
			nice := true
			for i := 0; i < 26; i++ {
				if (lowercase[i] && !uppercase[i]) || (uppercase[i] && !lowercase[i]) {
					nice = false
					break
				}
			}

			// Update longest nice substring if the current substring is nice and longer
			if nice && right-left+1 > len(longest) {
				longest = s[left : right+1]
			}
		}
	}

	return longest
}
