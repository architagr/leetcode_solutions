package reverse_string

func ReverseString(s []byte) {
	l := 0
	r := len(s) - 1
	for l <= r {
		s[l], s[r] = s[r], s[l]
		l++
		r--
	}
}
