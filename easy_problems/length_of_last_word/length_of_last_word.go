package length_of_last_word

func LengthOfLastWord(s string) int {
	count := 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == ' ' && count != 0 {
			break
		}
		if s[i] != ' ' {
			count++
		}
	}
	return count
}
