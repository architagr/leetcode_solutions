package numberofsubstringscontainingallthreecharacters

func numberOfSubstrings(s string) int {
	counts := make([]int, 3)
	result := 0
	left := 0

	for right := 0; right < len(s); right++ {
		counts[s[right]-'a']++

		for counts[0] > 0 && counts[1] > 0 && counts[2] > 0 {
			// All substrings starting from left to right are valid
			result += len(s) - right
			counts[s[left]-'a']--
			left++
		}
	}

	return result
}
