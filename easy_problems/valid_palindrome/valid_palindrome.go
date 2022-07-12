package valid_palindrome

func CheckPalindrome(s string) bool {
	n := len(s)
	j := n - 1
	if n <= 1 {
		return true
	}
	for i := 0; i <= j; i, j = i+1, j-1 {
		for i < n && checkInvalidChar(s[i]) {
			i++
		}

		for j >= 0 && checkInvalidChar(s[j]) {
			j--
		}

		if i < n && j >= 0 && checkCharSame_CaseInsensitive(s[i], s[j]) {
			return false
		}
	}
	return true
}

func checkInvalidChar(char byte) bool {
	return (char < 'A' || char > 'Z') && (char < 'a' || char > 'z') && (char < '0' || char > '9')
}

func checkCharSame_CaseInsensitive(char1, char2 byte) bool {
	return char1 != char2 && char1^32 != char2
}
