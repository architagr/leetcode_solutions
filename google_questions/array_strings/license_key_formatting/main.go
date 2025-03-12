package licensekeyformatting

import "strings"

func licenseKeyFormatting(s string, k int) string {
	parsedPart := []string{}

	x := 0
	temp := make([]byte, k)
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '-' {
			continue
		}
		c := s[i]

		if s[i] >= 'a' && s[i] <= 'z' {
			c = c & ^byte(32)
		}
		temp[x] = c
		x++
		if x == k {
			parsedPart = append(parsedPart, string(temp))
			x = 0
			temp = make([]byte, k)
		}
	}
	if x > 0 {
		parsedPart = append(parsedPart, string(temp[:x]))
	}

	return reverseString([]byte(strings.Join(parsedPart, "-")))
}

func reverseString(s []byte) string {
	l, r := 0, len(s)-1

	for l < r {
		s[l], s[r] = s[r], s[l]
		l++
		r--
	}
	return string(s)
}
