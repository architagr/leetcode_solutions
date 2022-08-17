package unique_morse_code_words

func UniqueMorseRepresentations(words []string) int {
	m := make(map[string]bool)
	for _, word := range words {
		m[convertWordToCode(word)] = true
	}
	return len(m)
}
func convertWordToCode(word string) string {
	codes := []string{".-", "-...", "-.-.", "-..", ".", "..-.", "--.", "....", "..", ".---", "-.-", ".-..", "--", "-.", "---", ".--.", "--.-", ".-.", "...", "-", "..-", "...-", ".--", "-..-", "-.--", "--.."}
	result := ""
	for i := 0; i < len(word); i++ {
		result += codes[int(word[i]-'a')]
	}
	return result
}
