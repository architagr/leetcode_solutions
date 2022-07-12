package implement_strstr

func StrStrApproch1(haystack, needle string) int {
	n, m := len(haystack), len(needle)
	if m > n {
		return -1
	}
	index := -1
	for i := 0; i < n; i++ {
		k := 0
		for i+k < n && k < m && haystack[i+k] == needle[k] {
			k++
			if k == m {
				return i
			}
		}
	}
	return index
}

func StrStrApproch2(haystack string, needle string) int {
	for i := 0; i <= len(haystack)-len(needle); i++ {
		if haystack[i:i+len(needle)] == needle {
			return i
		}
	}
	return -1
}
