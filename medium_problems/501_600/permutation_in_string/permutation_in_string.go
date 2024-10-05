package permutationinstring

func checkInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}
	s1Map := make(map[byte]int)
	subStringMap := make(map[byte]int)

	for i := 0; i < len(s1); i++ {
		s1Map[s1[i]]++
		subStringMap[s2[i]]++
	}
	if check(s1Map, subStringMap) {
		return true
	}
	for start, end := 1, len(s1); end < len(s2); start, end = start+1, end+1 {
		subStringMap[s2[start-1]]--
		if c := subStringMap[s2[start-1]]; c == 0 {
			delete(subStringMap, s2[start-1])
		}
		subStringMap[s2[end]]++
		if check(s1Map, subStringMap) {
			return true
		}
	}
	return false
}

func check(s1Map, subStringMap map[byte]int) bool {
	if len(s1Map) != len(subStringMap) {
		return false
	}
	for key, count := range s1Map {
		if c, ok := subStringMap[key]; !ok || (c > 0 && c != count) {
			return false
		}
	}
	return true
}
