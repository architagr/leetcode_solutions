package reverse_string_ii

func ReverseStr(s string, k int) string {
	ans := []byte(s)
	n := len(ans)
	for i := 0; i < n; i++ {
		l := i
		r := i + k
		if r >= n {
			r = n
		}

		ans = reversString(ans, l, r-1)
		i += k + k - 1
	}
	return string(ans)
}

func reversString(s []byte, l, r int) []byte {
	for l <= r {
		s[l], s[r] = s[r], s[l]
		l++
		r--
	}
	return s
}
