package reverse_words_in_a_string_iii

func ReverseWords(s string) string {
	s = s + " "
	ans := ""
	var space byte = ' '

	for i := 0; i < len(s); i++ {
		l := i
		for s[i] != space {
			i++
		}
		ans = ans + " " + string(reversString([]byte(s[l:i])))
	}
	return ans[1:]
}

func reversString(s []byte) []byte {
	l := 0
	r := len(s) - 1
	for l <= r {
		s[l], s[r] = s[r], s[l]
		l++
		r--
	}
	return s
}
