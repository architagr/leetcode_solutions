package count_number_of_consistent_strings

func CountConsistentStrings(allowed string, words []string) int {
	mapAllowed := make(map[byte]bool)
	count := 0
	for i := 0; i < len(allowed); i++ {
		mapAllowed[allowed[i]] = true
	}
	for _, word := range words {
		flag := true
		for i := 0; i < len(word); i++ {
			if !mapAllowed[word[i]] {
				flag = false
				break
			}
		}
		if flag {
			count++
		}
	}
	return count
}
