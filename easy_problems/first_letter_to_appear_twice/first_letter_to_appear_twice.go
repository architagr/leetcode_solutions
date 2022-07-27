package first_letter_to_appear_twice

func RepeatedCharacter(s string) byte {
	mapS := make(map[byte]int)
	ans := byte('a')
	for i := 0; i < len(s); i++ {
		if _, ok := mapS[s[i]]; ok {
			ans = s[i]
			break
		}
		mapS[s[i]]++
	}
	return ans
}
