package isomorphic_strings

func IsIsomorphic(s string, t string) bool {
	mapS := make(map[byte]byte)
	mapT := make(map[byte]bool)
	n := len(s)
	for i := 0; i < n; i++ {
		if val, ok := mapS[s[i]]; ok {
			if t[i] != val {
				return false
			}

		} else {
			if _, ok := mapT[t[i]]; ok {
				return false
			} else {
				mapS[s[i]] = t[i]
				mapT[t[i]] = true
			}
		}
	}
	return true
}
