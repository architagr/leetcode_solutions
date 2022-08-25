package ransom_note

func CanConstruct(ransomNote string, magazine string) bool {
	var counts [26]int

	if len(ransomNote) > len(magazine) {
		return false
	}

	for _, c := range magazine {
		counts[c-'a']++
	}

	for _, c := range ransomNote {
		if counts[c-'a'] == 0 {
			return false
		} else {
			counts[c-'a']--
		}
	}
	return true
}
