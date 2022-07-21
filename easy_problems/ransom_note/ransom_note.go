package ransom_note

func CanConstruct(ransomNote string, magazine string) bool {
	n, m := len(ransomNote), len(magazine)
	if n > m {
		return false
	}
	mapMag := make(map[byte]int)
	for i := 0; i < m; i++ {
		mapMag[magazine[i]]++
	}
	for i := 0; i < n; i++ {
		if mapMag[ransomNote[i]] > 0 {
			mapMag[ransomNote[i]]--
		} else {
			return false
		}
	}
	return true
}
