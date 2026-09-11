package findthedifference

func findTheDifference(s string, t string) byte {

	frequencyMapS := findChar(s)
	frequencyMapT := findChar(t)
	for i, count := range frequencyMapT {
		if frequencyMapS[i] < count {
			return byte('a' + i)
		}
	}
	return 'a'
}

func findChar(s string) []int {
	frequencyMap := make([]int, 26)
	for i := 0; i < len(s); i++ {
		frequencyMap[s[i]-'a']++
	}
	return frequencyMap
}
