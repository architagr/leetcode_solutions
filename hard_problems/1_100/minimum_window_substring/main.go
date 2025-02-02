package minimumwindowsubstring

func minWindow(s string, t string) string {
	frequencyMapT := make(map[byte]int, 26)
	for i := 0; i < len(t); i++ {
		frequencyMapT[t[i]]++
	}
	result := ""

	for l, r := 0, 0; r < len(s); r++ {
		if _, ok := frequencyMapT[s[r]]; !ok {
			continue
		}
		frequencyMapT[s[r]]--
		for ; found(frequencyMapT) && l <= r; l++ {
			if len(result) == 0 || r-l+1 < len(result) {
				result = s[l : r+1]
			}
			if _, ok := frequencyMapT[s[l]]; !ok {
				continue
			}
			frequencyMapT[s[l]]++
		}
	}
	return result
}
func found(frequencyMapT map[byte]int) bool {
	for _, c := range frequencyMapT {
		if c > 0 {
			return false
		}
	}
	return true
}
