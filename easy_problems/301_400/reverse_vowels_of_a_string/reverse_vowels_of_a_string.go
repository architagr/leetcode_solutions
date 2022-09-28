package reverse_vowels_of_a_string

func ReverseVowels(s string) string {
	ans := []byte(s)
	l := 0
	r := len(s) - 1
	for l <= r {
		if !check(s[l]) {
			l++
		} else if !check(s[r]) {
			r--
		} else {
			ans[l], ans[r] = ans[r], ans[l]
			l++
			r--
		}
	}
	return string(ans)
}

func check(a byte) bool {
	return a == 'a' || a == 'e' || a == 'i' || a == 'o' || a == 'u' || a == 'A' || a == 'E' || a == 'I' || a == 'O' || a == 'U'
}
