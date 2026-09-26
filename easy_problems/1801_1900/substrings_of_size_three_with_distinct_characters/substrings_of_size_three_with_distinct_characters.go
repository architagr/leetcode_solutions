package substringsofsizethreewithdistinctcharacters

func countGoodSubstrings(s string) int {
	n := len(s)
	if n < 3 {
		return 0
	}
	count := 0
	// Letter -> count within the current window. Keys are deleted at zero,
	// so len(m) is the number of distinct letters in the window.
	m := make(map[byte]int)
	m[s[0]]++
	m[s[1]]++
	m[s[2]]++
	if len(m) == 3 {
		count++
	}
	for i := 3; i < n; i++ {
		m[s[i-3]]--
		// Without this delete, a letter that has left the window would still
		// be a key, and len(m) would overcount.
		if m[s[i-3]] == 0 {
			delete(m, s[i-3])
		}
		m[s[i]]++
		if len(m) == 3 {
			count++
		}
	}
	return count
}
