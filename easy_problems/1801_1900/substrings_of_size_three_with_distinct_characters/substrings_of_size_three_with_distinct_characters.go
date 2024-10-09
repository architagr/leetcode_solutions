package substringsofsizethreewithdistinctcharacters

func countGoodSubstrings(s string) int {
	n := len(s)
	if n < 3 {
		return 0
	}
	count := 0
	m := make(map[byte]int)
	m[s[0]]++
	m[s[1]]++
	m[s[2]]++
	if len(m) == 3 {
		count++
	}
	for i := 3; i < n; i++ {
		m[s[i-3]]--
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
