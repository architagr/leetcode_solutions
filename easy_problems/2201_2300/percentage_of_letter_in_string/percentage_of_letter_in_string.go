package percentage_of_letter_in_string

func percentageLetter(s string, letter byte) int {
	count := 0
	for i := 0; i < len(s); i++ {
		if s[i] == letter {
			count++
		}
	}
	return (count * 100) / len(s)
}
