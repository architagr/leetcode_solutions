package numberofsubstringscontainingallthreecharacters

func numberOfSubstrings(s string) int {
	counts := make([]int, 3)
	result := 0
	left := 0

	for right := 0; right < len(s); right++ {
		counts[s[right]-'a']++

		for counts[0] > 0 && counts[1] > 0 && counts[2] > 0 {
			// Adding letters can't remove a, b or c, so every substring that
			// starts at left and ends at right or later is valid.
			result += len(s) - right
			counts[s[left]-'a']--
			left++
		}
	}

	return result
}
