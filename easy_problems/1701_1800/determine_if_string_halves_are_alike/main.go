package determineifstringhalvesarealike

func halvesAreAlike(s string) bool {
	counta, countb := 0, 0
	string1 := s[:len(s)/2]
	string2 := s[len(s)/2:]

	for i := 0; i < len(string1); i++ {
		if check(string1[i]) {
			counta++
		}
		if check(string2[i]) {
			countb++
		}
	}
	return counta == countb
}
func check(ch byte) bool {
	return ch == 'a' || ch == 'e' || ch == 'i' || ch == 'o' || ch == 'u' || ch == 'A' || ch == 'E' || ch == 'I' || ch == 'O' || ch == 'U'
}
