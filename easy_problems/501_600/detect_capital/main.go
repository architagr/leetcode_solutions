package detectcapital

func detectCapitalUse(word string) bool {
	n := len(word)
	count := 0

	for i := 0; i < n; i++ {
		if word[i] >= 'A' && word[i] <= 'Z' {
			count++
			if count < i+1 {
				return false
			}
		}
	}

	return count == n || count == 1 || count == 0

}
