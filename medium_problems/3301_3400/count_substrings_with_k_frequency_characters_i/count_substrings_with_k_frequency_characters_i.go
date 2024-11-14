package countsubstringswithkfrequencycharactersi

func numberOfSubstrings(s string, k int) int {
	counter := make([]int, 26)
	n := len(s)
	count := 0
	l := 0

	for left, right := 0, 0; right < n; right++ {
		counter[s[right]-'a']++
		for counter[s[right]-'a'] >= k {
			counter[s[left]-'a']--
			left++
			l++
		}
		count += l
	}
	return count
}
