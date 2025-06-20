package longestpalindrome

func longestPalindrome(s string) int {
	charMap := make(map[byte]int, 52)

	for i := 0; i < len(s); i++ {
		charMap[s[i]]++
	}
	result := 0
	odd := 0

	for _, v := range charMap {
		if v&1 == 1 {
			if odd == 0 {
				odd = 1
			}
			v -= 1
		}
		result += v

	}
	return result + odd
}
